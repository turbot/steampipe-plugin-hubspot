package hubspot

import (
	"context"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHubSpotContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_contact",
		Description: "ACL entries for the service version.",
		List: &plugin.ListConfig{
			Hydrate: listContacts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "acl_id",
					Require: plugin.Optional,
				},
			},
		},
		// Get: &plugin.GetConfig{
		// 	Hydrate:    getACLEntry,
		// 	KeyColumns: plugin.AllColumns([]string{"acl_id", "id"}),
		// },
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ACL entry.",
			},
			{
				Name:        "acl_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying a ACL.",
			},
			{
				Name:        "ip",
				Type:        proto.ColumnType_IPADDR,
				Description: "An IP address.",
			},
			{
				Name:        "negated",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was deleted.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was updated.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listContacts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_acl_entry.listACLEntries", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := contacts.NewAPIClient(contacts.NewConfiguration())
	response, _, err := client.BasicApi.GetPage(context).Execute()
	if err != nil {
		return nil, err
	}
	// // Limiting the results
	// maxLimit := 1000
	// if d.QueryContext.Limit != nil {
	// 	limit := int(*d.QueryContext.Limit)
	// 	if limit < maxLimit {
	// 		maxLimit = limit
	// 	}
	// }
	for _, test := range response.Results {
		d.StreamListItem(ctx, test)
	}

	return nil, nil
}

// func getACLEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	aclId := d.EqualsQuals["acl_id"].GetStringValue()
// 	id := d.EqualsQuals["id"].GetStringValue()

// 	// check if aclId or id is empty
// 	if aclId == "" || id == "" {
// 		return nil, nil
// 	}

// 	serviceClient, err := connect(ctx, d)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("hubspot_acl_entry.getACLEntry", "connection_error", err)
// 		return nil, err
// 	}

// 	input := &hubspot.GetACLEntryInput{
// 		ServiceID: serviceClient.ServiceID,
// 		ACLID:     aclId,
// 		ID:        id,
// 	}
// 	result, err := serviceClient.Client.GetACLEntry(input)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("hubspot_acl_entry.getACLEntry", "api_error", err)
// 		return nil, err
// 	}

// 	return result, nil
// }
