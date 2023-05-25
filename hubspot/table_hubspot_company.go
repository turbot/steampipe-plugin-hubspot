package hubspot

import (
	"context"
	"time"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/companies"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHubSpotCompany(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_company",
		Description: "List of HubSpot Companies.",
		List: &plugin.ListConfig{
			Hydrate: listCompanies,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "archived",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCompany,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the company.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp when the company was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp when the company was last updated.",
			},
			{
				Name:        "archived",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the company is archived or not.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_STRING,
				Description: "The timestamp when the company was archived.",
			},
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Description: "The domain associated with the company.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the company.",
			},
			{
				Name:        "properties",
				Type:        proto.ColumnType_JSON,
				Description: "The properties associated with the company.",
				Hydrate:     getCompanyProperties,
				Transform:   transform.FromField("Properties"),
			},
			{
				Name:        "properties_with_history",
				Type:        proto.ColumnType_JSON,
				Description: "The properties associated with the company including historical changes.",
				Hydrate:     getCompanyProperties,
				Transform:   transform.FromField("PropertiesWithHistory"),
			},
			{
				Name:        "associations_with_contacts",
				Type:        proto.ColumnType_JSON,
				Description: "The associations of the company with contacts.",
				Hydrate:     getCompanyAssociationsWithContacts,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "associations_with_deals",
				Type:        proto.ColumnType_JSON,
				Description: "The associations of the company with deals.",
				Hydrate:     getCompanyAssociationsWithDeals,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "associations_with_tickets",
				Type:        proto.ColumnType_JSON,
				Description: "The associations of the company with tickets.",
				Hydrate:     getCompanyAssociationsWithTickets,
				Transform:   transform.FromValue(),
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

type Company struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Archived   *bool
	ArchivedAt *time.Time
	Domain     string
	Name       string
}

func listCompanies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.listCompanies", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := companies.NewAPIClient(companies.NewConfiguration())

	// Limiting the results
	var maxLimit int32 = 100
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}
	var after string = ""
	archived := false

	if d.EqualsQuals["archived"] != nil {
		archived = d.EqualsQuals["archived"].GetBoolValue()
	}

	for {
		if after == "" {
			response, _, err := client.BasicApi.GetPage(context).Limit(maxLimit).Archived(archived).Execute()
			if err != nil {
				plugin.Logger(ctx).Error("hubspot_company.listCompanies", "api_error", err)
				return nil, err
			}
			for _, company := range response.Results {
				d.StreamListItem(ctx, Company{company.Id, company.CreatedAt, company.UpdatedAt, company.Archived, company.ArchivedAt, company.Properties["domain"], company.Properties["name"]})

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
			if !response.Paging.HasNext() {
				break
			}
			after = response.Paging.Next.After
		} else {
			response, _, err := client.BasicApi.GetPage(context).Limit(maxLimit).After(after).Archived(archived).Execute()
			if err != nil {
				plugin.Logger(ctx).Error("hubspot_company.listCompanies", "api_error", err)
				return nil, err
			}
			for _, company := range response.Results {
				d.StreamListItem(ctx, Company{company.Id, company.CreatedAt, company.UpdatedAt, company.Archived, company.ArchivedAt, company.Properties["domain"], company.Properties["name"]})

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
			if !response.Paging.HasNext() {
				break
			}
			after = response.Paging.Next.After
		}
	}

	return nil, nil
}

func getCompany(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompany", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := companies.NewAPIClient(companies.NewConfiguration())

	company, _, err := client.BasicApi.GetByID(context, id).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompany", "api_error", err)
		return nil, err
	}

	return Company{company.Id, company.CreatedAt, company.UpdatedAt, company.Archived, company.ArchivedAt, company.Properties["domain"], company.Properties["name"]}, nil
}

func getCompanyProperties(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Company).Id

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompanyProperties", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := companies.NewAPIClient(companies.NewConfiguration())
	properties, err := listAllPropertiesByObjectType(ctx, d, "company")
	if err != nil {
		return nil, err
	}

	company, _, err := client.BasicApi.GetByID(context, id).PropertiesWithHistory(properties).Properties(properties).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompanyProperties", "api_error", err)
		return nil, err
	}

	return company, nil
}

func getCompanyAssociationsWithContacts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Company).Id

	associatedIds, err := getAssociations(ctx, d, id, "company", "contact")
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompanyAssociationsWithContacts", "api_error", err)
		return nil, err
	}

	return associatedIds, nil
}

func getCompanyAssociationsWithDeals(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Company).Id

	associatedIds, err := getAssociations(ctx, d, id, "company", "deal")
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompanyAssociationsWithDeals", "api_error", err)
		return nil, err
	}

	return associatedIds, nil
}

func getCompanyAssociationsWithTickets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Company).Id

	associatedIds, err := getAssociations(ctx, d, id, "company", "ticket")
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getCompanyAssociationsWithTickets", "api_error", err)
		return nil, err
	}

	return associatedIds, nil
}
