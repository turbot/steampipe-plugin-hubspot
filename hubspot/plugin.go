package hubspot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-hubspot",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
		},
		TableMap: map[string]*plugin.Table{
			"hubspot_blog_post": tableHubSpotBlogPost(ctx),
			"hubspot_company":   tableHubSpotCompany(ctx),
			"hubspot_contact":   tableHubSpotContact(ctx),
			"hubspot_deal":      tableHubSpotDeal(ctx),
			"hubspot_domain":    tableHubSpotDomain(ctx),
			"hubspot_hub_db":    tableHubSpotHubDB(ctx),
			"hubspot_owner":     tableHubSpotOwner(ctx),
			"hubspot_ticket":    tableHubSpotTicket(ctx),
		},
	}
	return p
}
