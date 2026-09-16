package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

type hubspotAccessToken struct {
	UserId      int64    `json:"userId"`
	HubId       int64    `json:"hubId"`
	AppId       int64    `json:"appId"`
	IsUserToken bool     `json:"isUserToken"`
	Scopes      []string `json:"scopes"`
}

//// TABLE DEFINITION

func tableHubSpotAccessToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_access_token",
		Description: "Details about the private app access token in use, including its granted scopes.",
		List: &plugin.ListConfig{
			Hydrate: listAccessToken,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "hub_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the HubSpot account the token belongs to.",
			},
			{
				Name:        "app_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the private app the token belongs to.",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the user the token is associated with.",
			},
			{
				Name:        "is_user_token",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the token is a user access token.",
			},
			{
				Name:        "scopes",
				Type:        proto.ColumnType_JSON,
				Description: "The scopes granted to the access token.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AppId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listAccessToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_access_token.listAccessToken", "connection_error", err)
		return nil, err
	}

	body := map[string]string{"tokenKey": authorizer.Token}

	var token hubspotAccessToken
	if err := hubspotPost(ctx, d, "/oauth/v2/private-apps/get/access-token-info", body, &token); err != nil {
		// The introspection endpoint is not served on every account/region and
		// returns 404 when unavailable. Treat that as no rows so it never breaks
		// the rest of the plugin, but surface any other failure.
		if shouldIgnoreErrors([]string{"404"})(ctx, d, h, err) {
			plugin.Logger(ctx).Warn("hubspot_access_token.listAccessToken", "token_unavailable", err)
			return nil, nil
		}
		plugin.Logger(ctx).Error("hubspot_access_token.listAccessToken", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, token)

	return nil, nil
}
