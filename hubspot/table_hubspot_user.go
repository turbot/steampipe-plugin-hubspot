package hubspot

import (
	"context"
	"net/url"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type hubspotUser struct {
	Id               string   `json:"id"`
	Email            string   `json:"email"`
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	RoleId           string   `json:"roleId"`
	RoleIds          []string `json:"roleIds"`
	PrimaryTeamId    string   `json:"primaryTeamId"`
	SecondaryTeamIds []string `json:"secondaryTeamIds"`
	SuperAdmin       bool     `json:"superAdmin"`
}

type hubspotUserListResponse struct {
	Results []hubspotUser  `json:"results"`
	Paging  *hubspotPaging `json:"paging"`
}

//// TABLE DEFINITION

func tableHubSpotUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_user",
		Description: "List of HubSpot account users, including their roles, teams, and super admin status.",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the user.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address of the user.",
			},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Description: "The first name of the user.",
			},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Description: "The last name of the user.",
			},
			{
				Name:        "super_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the user is a super admin.",
			},
			{
				Name:        "role_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the primary role assigned to the user.",
			},
			{
				Name:        "primary_team_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the user's primary team.",
			},
			{
				Name:        "role_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The IDs of all roles assigned to the user.",
			},
			{
				Name:        "secondary_team_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The IDs of the user's secondary teams.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Email"),
			},
		}),
	}
}

//// LIST FUNCTION

func listUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		q := url.Values{}
		q.Set("limit", strconv.Itoa(int(maxLimit)))
		if after != "" {
			q.Set("after", after)
		}

		var response hubspotUserListResponse
		if err := hubspotGet(ctx, d, "/settings/v3/users?"+q.Encode(), &response); err != nil {
			plugin.Logger(ctx).Error("hubspot_user.listUsers", "api_error", err)
			return nil, err
		}

		for _, user := range response.Results {
			d.StreamListItem(ctx, user)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.Paging == nil || response.Paging.Next == nil || response.Paging.Next.After == "" {
			break
		}
		after = response.Paging.Next.After
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")
	if id == "" {
		return nil, nil
	}

	var user hubspotUser
	if err := hubspotGet(ctx, d, "/settings/v3/users/"+url.PathEscape(id), &user); err != nil {
		plugin.Logger(ctx).Error("hubspot_user.getUser", "api_error", err)
		return nil, err
	}

	return user, nil
}
