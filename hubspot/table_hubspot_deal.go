package hubspot

import (
	"context"
	"time"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/deals"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHubSpotDeal(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_deal",
		Description: "List of HubSpot Deals.",
		List: &plugin.ListConfig{
			Hydrate: listDeals,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "archived",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getDeal,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "Updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "archived",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "amount",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "deal_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "pipeline",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "close_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "deal_stage",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "properties",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getDealProperties,
				Transform:   transform.FromField("Properties"),
			},
			{
				Name:        "properties_with_history",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getDealProperties,
				Transform:   transform.FromField("PropertiesWithHistory"),
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DealName"),
			},
		},
	}
}

type Deal struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Archived   *bool
	ArchivedAt *time.Time
	Amount     string
	DealName   string
	Pipeline   string
	CloseDate  string
	DealStage  string
}

func listDeals(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_deal.listDeals", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := deals.NewAPIClient(deals.NewConfiguration())

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
				plugin.Logger(ctx).Error("hubspot_deal.listDeals", "api_error", err)
				return nil, err
			}
			for _, deal := range response.Results {
				d.StreamListItem(ctx, Deal{deal.Id, deal.CreatedAt, deal.UpdatedAt, deal.Archived, deal.ArchivedAt, deal.Properties["amount"], deal.Properties["dealname"], deal.Properties["pipeline"], deal.Properties["closedate"], deal.Properties["dealstage"]})

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
				plugin.Logger(ctx).Error("hubspot_deal.listDeals", "api_error", err)
				return nil, err
			}
			for _, deal := range response.Results {
				d.StreamListItem(ctx, Deal{deal.Id, deal.CreatedAt, deal.UpdatedAt, deal.Archived, deal.ArchivedAt, deal.Properties["amount"], deal.Properties["dealname"], deal.Properties["pipeline"], deal.Properties["closedate"], deal.Properties["dealstage"]})

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

func getDeal(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_deal.getDeal", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := deals.NewAPIClient(deals.NewConfiguration())

	deal, _, err := client.BasicApi.GetByID(context, id).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_deal.getDeal", "api_error", err)
		return nil, err
	}

	return Deal{deal.Id, deal.CreatedAt, deal.UpdatedAt, deal.Archived, deal.ArchivedAt, deal.Properties["amount"], deal.Properties["dealname"], deal.Properties["pipeline"], deal.Properties["closedate"], deal.Properties["dealstage"]}, nil
}

func getDealProperties(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Deal).Id

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_deal.getDealProperties", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := deals.NewAPIClient(deals.NewConfiguration())
	properties, err := listAllPropertiesByObjectType(ctx, d, "deal")
	if err != nil {
		return nil, err
	}

	deal, _, err := client.BasicApi.GetByID(context, id).PropertiesWithHistory(properties).Properties(properties).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_deal.getDealProperties", "api_error", err)
		return nil, err
	}

	return deal, nil
}
