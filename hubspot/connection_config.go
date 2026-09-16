package hubspot

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type hubSpotConfig struct {
	PrivateAppToken *string `hcl:"private_app_token"`
}

func ConfigInstance() interface{} {
	return &hubSpotConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hubSpotConfig {
	if connection == nil || connection.GetConfig() == nil {
		return hubSpotConfig{}
	}
	config, _ := connection.GetConfig().(hubSpotConfig)
	return config
}
