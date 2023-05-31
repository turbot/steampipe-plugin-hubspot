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
				Description: "The unique ID of the deal.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp when the deal was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp when the deal was last updated.",
			},
			{
				Name:        "archived",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the deal is archived or not.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_STRING,
				Description: "The timestamp when the deal was archived.",
			},
			{
				Name:        "amount",
				Type:        proto.ColumnType_STRING,
				Description: "The amount associated with the deal.",
			},
			{
				Name:        "deal_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the deal.",
			},
			{
				Name:        "pipeline",
				Type:        proto.ColumnType_STRING,
				Description: "The pipeline associated with the deal.",
			},
			{
				Name:        "close_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The expected close date of the deal.",
			},
			{
				Name:        "deal_stage",
				Type:        proto.ColumnType_STRING,
				Description: "The stage of the deal.",
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
