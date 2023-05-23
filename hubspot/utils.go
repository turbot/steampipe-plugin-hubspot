package hubspot

import (
	"context"
	"os"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*hubspot.TokenAuthorizer, error) {

	// Default to the env var settings
	token := os.Getenv("FASTLY_API_KEY")
	// baseURL := os.Getenv("FASTLY_API_URL")
	// serviceID := os.Getenv("FASTLY_SERVICE_ID")
	// serviceVersion := os.Getenv("FASTLY_SERVICE_VERSION")

	// Prefer config settings
	hubSpotConfig := GetConfig(d.Connection)
	if hubSpotConfig.Token != nil {
		token = *hubSpotConfig.Token
	}

	authorizer := hubspot.NewTokenAuthorizer(token)
	//ctx := hubspot.WithAuthorizer(context.Background(), authorizer)

	return authorizer, nil
}
