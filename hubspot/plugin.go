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
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

func pluginTableDefinitions(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {

	// set Connection and ConectionCache
	queryData := &plugin.QueryData{
		Connection:      d.Connection,
		ConnectionCache: d.ConnectionCache,
	}

	companyPropertiesColumns, err := listAllPropertiesByObjectType(ctx, queryData, "company")
	if err != nil {
		return nil, err
	}

	// Initialize tables
	tables := map[string]*plugin.Table{
		"hubspot_blog_post": tableHubSpotBlogPost(ctx),
		"hubspot_company":   tableHubSpotCompany(ctx, companyPropertiesColumns),
		"hubspot_contact":   tableHubSpotContact(ctx),
		"hubspot_deal":      tableHubSpotDeal(ctx),
		"hubspot_domain":    tableHubSpotDomain(ctx),
		"hubspot_hub_db":    tableHubSpotHubDB(ctx),
		"hubspot_owner":     tableHubSpotOwner(ctx),
		"hubspot_ticket":    tableHubSpotTicket(ctx),
	}

	return tables, nil
}
