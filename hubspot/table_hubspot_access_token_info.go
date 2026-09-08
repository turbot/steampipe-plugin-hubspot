package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type hubspotAccessTokenInfo struct {
	UserId      int64    `json:"userId"`
	HubId       int64    `json:"hubId"`
	AppId       int64    `json:"appId"`
	IsUserToken bool     `json:"isUserToken"`
	Scopes      []string `json:"scopes"`
}

//// TABLE DEFINITION

func tableHubSpotAccessTokenInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_access_token_info",
		Description: "Details about the private app access token in use, including its granted scopes.",
		List: &plugin.ListConfig{
			Hydrate: listAccessTokenInfo,
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

func listAccessTokenInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_access_token_info.listAccessTokenInfo", "connection_error", err)
		return nil, err
	}

	body := map[string]string{"tokenKey": authorizer.Token}

	var info hubspotAccessTokenInfo
	if err := hubspotPost(ctx, d, "/oauth/v2/private-apps/get/access-token-info", body, &info); err != nil {
		// The introspection endpoint is not served on every account/region and
		// returns 404 when unavailable. Treat that as no rows so it never breaks
		// the rest of the plugin, but surface any other failure.
		if shouldIgnoreErrors([]string{"404"})(ctx, d, h, err) {
			plugin.Logger(ctx).Warn("hubspot_access_token_info.listAccessTokenInfo", "token_info_unavailable", err)
			return nil, nil
		}
		plugin.Logger(ctx).Error("hubspot_access_token_info.listAccessTokenInfo", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, info)

	return nil, nil
}
