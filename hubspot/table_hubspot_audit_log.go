package hubspot

import (
	"context"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/audit_logs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableHubSpotAuditLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_audit_log",
		Description: "History of CMS content changes (create, update, publish, delete) in the HubSpot account.",
		List: &plugin.ListConfig{
			Hydrate: listAuditLogs,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Optional,
				},
				{
					Name:    "object_type",
					Require: plugin.Optional,
				},
				{
					Name:    "event",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "object_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the object that was changed.",
				Transform:   transform.FromField("ObjectId"),
			},
			{
				Name:        "object_name",
				Type:        proto.ColumnType_STRING,
				Description: "The internal name of the object in HubSpot.",
			},
			{
				Name:        "object_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.).",
			},
			{
				Name:        "event",
				Type:        proto.ColumnType_STRING,
				Description: "The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED, RESTORE).",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the user who caused the event.",
				Transform:   transform.FromField("UserId"),
			},
			{
				Name:        "full_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the user who caused the event.",
			},
			{
				Name:        "timestamp",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp at which the event occurred.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ObjectId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_audit_log.listAuditLogs", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := audit_logs.NewAPIClient(audit_logs.NewConfiguration())

	// Limiting the results
	var maxLimit int32 = 100
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	after := ""
	for {
		req := client.AuditLogsApi.GetPage(context).Limit(maxLimit)
		if after != "" {
			req = req.After(after)
		}
		if q := d.EqualsQualString("user_id"); q != "" {
			req = req.UserId([]string{q})
		}
		if q := d.EqualsQualString("object_type"); q != "" {
			req = req.ObjectType([]string{q})
		}
		if q := d.EqualsQualString("event"); q != "" {
			req = req.EventType([]string{q})
		}

		response, _, err := req.Execute()
		if err != nil {
			plugin.Logger(ctx).Error("hubspot_audit_log.listAuditLogs", "api_error", err)
			return nil, err
		}

		for _, auditLog := range response.Results {
			d.StreamListItem(ctx, auditLog)

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

	return nil, nil
}
