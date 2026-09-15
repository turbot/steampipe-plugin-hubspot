package hubspot

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type hubspotLoginActivity struct {
	Id             string    `json:"id"`
	LoginAt        time.Time `json:"loginAt"`
	UserId         int64     `json:"userId"`
	Email          string    `json:"email"`
	LoginSucceeded bool      `json:"loginSucceeded"`
	IpAddress      string    `json:"ipAddress"`
	Location       string    `json:"location"`
	UserAgent      string    `json:"userAgent"`
	CountryCode    string    `json:"countryCode"`
	RegionCode     string    `json:"regionCode"`
}

type hubspotLoginActivityListResponse struct {
	Results []hubspotLoginActivity `json:"results"`
	Paging  *hubspotPaging         `json:"paging"`
}

//// TABLE DEFINITION

func tableHubSpotLoginActivity(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_login_activity",
		Description: "History of user login attempts (successful and unsuccessful) in the HubSpot account over the past 90 days.",
		List: &plugin.ListConfig{
			Hydrate: listLoginActivities,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the login event.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "login_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when login was attempted.",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the user associated with the activity.",
				Transform:   transform.FromField("UserId"),
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address of the user associated with the activity.",
			},
			{
				Name:        "login_succeeded",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the login attempt was successful.",
			},
			{
				Name:        "ip_address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address used for the login attempt.",
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "The location where the login was attempted.",
			},
			{
				Name:        "user_agent",
				Type:        proto.ColumnType_STRING,
				Description: "User agent information about the device used for login.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "The country code of the location where the login was attempted.",
			},
			{
				Name:        "region_code",
				Type:        proto.ColumnType_STRING,
				Description: "The region code of the location where the login was attempted.",
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

func listLoginActivities(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		if d.EqualsQuals["user_id"] != nil {
			q.Set("userId", strconv.FormatInt(d.EqualsQuals["user_id"].GetInt64Value(), 10))
		}

		var response hubspotLoginActivityListResponse
		if err := hubspotGet(ctx, d, "/account-info/v3/activity/login?"+q.Encode(), &response); err != nil {
			plugin.Logger(ctx).Error("hubspot_login_activity.listLoginActivities", "api_error", err)
			return nil, err
		}

		for _, activity := range response.Results {
			d.StreamListItem(ctx, activity)

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
