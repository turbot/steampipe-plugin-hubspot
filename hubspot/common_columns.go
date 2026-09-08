package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "portal_id",
			Description: "Unique identifier for the HubSpot portal or account.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getPortalId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getPortalIdMemoized = plugin.HydrateFunc(getPortalInfoUncached).Memoize(memoize.WithCacheKeyFunction(getPortalIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getPortalId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	acc, err := getPortalIdMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return acc.(AccountInfo).PortalID, nil
}

// Build a cache key for the call to getPortalIdCacheKey.
func getPortalIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getPortalId"
	return key, nil
}

func getPortalInfoUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var accInfo AccountInfo
	if err := hubspotGet(ctx, d, "/account-info/v3/details", &accInfo); err != nil {
		plugin.Logger(ctx).Error("getPortalInfoUncached", "api_error", err)
		return nil, err
	}

	return accInfo, nil
}

type AccountInfo struct {
	PortalID              int64    `json:"portalId"`
	AccountType           string   `json:"accountType"`
	TimeZone              string   `json:"timeZone"`
	CompanyCurrency       string   `json:"companyCurrency"`
	AdditionalCurrencies  []string `json:"additionalCurrencies"`
	UTCOffset             string   `json:"utcOffset"`
	UTCOffsetMilliseconds int64    `json:"utcOffsetMilliseconds"`
	UIDomain              string   `json:"uiDomain"`
	DataHostingLocation   string   `json:"dataHostingLocation"`
}
