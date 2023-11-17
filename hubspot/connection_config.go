package hubspot

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type hubSpotConfig struct {
	PrivateAppToken *string `cty:"private_app_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"private_app_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &hubSpotConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hubSpotConfig {
	if connection == nil || connection.Config == nil {
		return hubSpotConfig{}
	}
	config, _ := connection.Config.(hubSpotConfig)
	return config
}
