package hubspot

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

type hubspotSecurityActivity struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UserId      int64     `json:"userId"`
	Type        string    `json:"type"`
	ActingUser  string    `json:"actingUser"`
	ObjectId    string    `json:"objectId"`
	InfoUrl     string    `json:"infoUrl"`
	Location    string    `json:"location"`
	IpAddress   string    `json:"ipAddress"`
	CountryCode string    `json:"countryCode"`
	RegionCode  string    `json:"regionCode"`
}

type hubspotSecurityActivityListResponse struct {
	Results []hubspotSecurityActivity `json:"results"`
	Paging  *hubspotPaging            `json:"paging"`
}

//// TABLE DEFINITION

func tableHubSpotSecurityActivity(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_security_activity",
		Description: "History of security-related user activity in the HubSpot account, such as adding admins, enabling SSO, and installing integrations.",
		List: &plugin.ListConfig{
			Hydrate: listSecurityActivities,
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
				Description: "The unique ID of the activity.",
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the activity occurred.",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the user associated with the activity.",
				Transform:   transform.FromField("UserId"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of activity (e.g. ADD_ADMIN_USER, ADD_SINGLE_SIGN_ON, INSTALL_INTEGRATION).",
			},
			{
				Name:        "acting_user",
				Type:        proto.ColumnType_STRING,
				Description: "The email address of the user associated with the activity.",
			},
			{
				Name:        "object_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the object that the activity was performed on.",
			},
			{
				Name:        "info_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the page where the activity occurred.",
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "The location where the activity took place.",
			},
			{
				Name:        "ip_address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address where the activity originated.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "The country code of the location where the activity took place.",
			},
			{
				Name:        "region_code",
				Type:        proto.ColumnType_STRING,
				Description: "The region code of the location where the activity took place.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type"),
			},
		}),
	}
}

//// LIST FUNCTION

func listSecurityActivities(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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

		var response hubspotSecurityActivityListResponse
		if err := hubspotGet(ctx, d, "/account-info/v3/activity/security?"+q.Encode(), &response); err != nil {
			plugin.Logger(ctx).Error("hubspot_security_activity.listSecurityActivities", "api_error", err)
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
