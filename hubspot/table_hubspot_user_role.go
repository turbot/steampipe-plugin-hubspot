package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type hubspotUserRole struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	RequiresBillingWrite bool   `json:"requiresBillingWrite"`
}

type hubspotUserRoleListResponse struct {
	Results []hubspotUserRole `json:"results"`
}

//// TABLE DEFINITION

func tableHubSpotUserRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_user_role",
		Description: "List of roles defined in the HubSpot account.",
		List: &plugin.ListConfig{
			Hydrate: listUserRoles,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the role.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the role.",
			},
			{
				Name:        "requires_billing_write",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the role requires billing write access.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listUserRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var response hubspotUserRoleListResponse
	if err := hubspotGet(ctx, d, "/settings/v3/users/roles", &response); err != nil {
		plugin.Logger(ctx).Error("hubspot_user_role.listUserRoles", "api_error", err)
		return nil, err
	}

	for _, role := range response.Results {
		d.StreamListItem(ctx, role)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
