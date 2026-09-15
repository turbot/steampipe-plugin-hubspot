package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type hubspotTeam struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	UserIds          []string `json:"userIds"`
	SecondaryUserIds []string `json:"secondaryUserIds"`
}

type hubspotTeamListResponse struct {
	Results []hubspotTeam `json:"results"`
}

//// TABLE DEFINITION

func tableHubSpotTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_team",
		Description: "List of teams defined in the HubSpot account, including their member users.",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the team.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the team.",
			},
			{
				Name:        "user_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The IDs of the users whose primary team is this team.",
			},
			{
				Name:        "secondary_user_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The IDs of the users who have this team as a secondary team.",
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

func listTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var response hubspotTeamListResponse
	if err := hubspotGet(ctx, d, "/settings/v3/users/teams", &response); err != nil {
		plugin.Logger(ctx).Error("hubspot_team.listTeams", "api_error", err)
		return nil, err
	}

	for _, team := range response.Results {
		d.StreamListItem(ctx, team)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
