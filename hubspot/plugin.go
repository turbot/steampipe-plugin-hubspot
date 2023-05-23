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
			"hubspot_contact": tableHubSpotContact(ctx),
			// "hubspot_acl_entry":       tableHubSpotACLEntry(ctx),
			// "hubspot_backend":         tableHubSpotBackend(ctx),
			// "hubspot_condition":       tableHubSpotCondition(ctx),
			// "hubspot_data_center":     tableHubSpotDataCenter(ctx),
			// "hubspot_dictionary":      tableHubSpotDictionary(ctx),
			// "hubspot_health_check":    tableHubSpotHealthCheck(ctx),
			// "hubspot_ip_range":        tableHubSpotIPRange(ctx),
			// "hubspot_pool":            tableHubSpotPool(ctx),
			// "hubspot_service":         tableHubSpotService(ctx),
			// "hubspot_service_domain":  tableHubSpotServiceDomain(ctx),
			// "hubspot_service_version": tableHubSpotServiceVersion(ctx),
			// "hubspot_token":           tableHubSpotToken(ctx),
		},
	}
	return p
}
