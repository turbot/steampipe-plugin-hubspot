package hubspot

import (
	"context"
	"time"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/tickets"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHubSpotTicket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_ticket",
		Description: "List of HubSpot Tickets.",
		List: &plugin.ListConfig{
			Hydrate: listTickets,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "archived",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getTicket,
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
				Name:        "content",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "subject",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "pipeline",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "pipeline_stage",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "ticket_category",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "ticket_priority",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "properties",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getTicketProperties,
				Transform:   transform.FromField("Properties"),
			},
			{
				Name:        "properties_with_history",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getTicketProperties,
				Transform:   transform.FromField("PropertiesWithHistory"),
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Subject"),
			},
		},
	}
}

type Ticket struct {
	Id             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Archived       *bool
	ArchivedAt     *time.Time
	Content        string
	Subject        string
	Pipeline       string
	PipelineStage  string
	TicketCategory string
	TicketPriority string
}

func listTickets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_ticket.listTickets", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := tickets.NewAPIClient(tickets.NewConfiguration())

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
				plugin.Logger(ctx).Error("hubspot_ticket.listTickets", "api_error", err)
				return nil, err
			}
			for _, ticket := range response.Results {
				d.StreamListItem(ctx, Ticket{ticket.Id, ticket.CreatedAt, ticket.UpdatedAt, ticket.Archived, ticket.ArchivedAt, ticket.Properties["content"], ticket.Properties["subject"], ticket.Properties["hs_pipeline"], ticket.Properties["hs_pipeline_stage"], ticket.Properties["hs_ticket_category"], ticket.Properties["hs_ticket_priority"]})

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
				plugin.Logger(ctx).Error("hubspot_ticket.listTickets", "api_error", err)
				return nil, err
			}
			for _, ticket := range response.Results {
				d.StreamListItem(ctx, Ticket{ticket.Id, ticket.CreatedAt, ticket.UpdatedAt, ticket.Archived, ticket.ArchivedAt, ticket.Properties["content"], ticket.Properties["subject"], ticket.Properties["hs_pipeline"], ticket.Properties["hs_pipeline_stage"], ticket.Properties["hs_ticket_category"], ticket.Properties["hs_ticket_priority"]})

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

func getTicket(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_ticket.getTicket", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := tickets.NewAPIClient(tickets.NewConfiguration())

	ticket, _, err := client.BasicApi.GetByID(context, id).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_ticket.getTicket", "api_error", err)
		return nil, err
	}

	return Ticket{ticket.Id, ticket.CreatedAt, ticket.UpdatedAt, ticket.Archived, ticket.ArchivedAt, ticket.Properties["content"], ticket.Properties["subject"], ticket.Properties["hs_pipeline"], ticket.Properties["hs_pipeline_stage"], ticket.Properties["hs_ticket_category"], ticket.Properties["hs_ticket_priority"]}, nil
}

func getTicketProperties(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(Ticket).Id

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_ticket.getTicketProperties", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := tickets.NewAPIClient(tickets.NewConfiguration())
	properties, err := listAllPropertiesByObjectType(ctx, d, "ticket")
	if err != nil {
		return nil, err
	}

	ticket, _, err := client.BasicApi.GetByID(context, id).PropertiesWithHistory(properties).Properties(properties).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_ticket.getTicketProperties", "api_error", err)
		return nil, err
	}

	return ticket, nil
}
