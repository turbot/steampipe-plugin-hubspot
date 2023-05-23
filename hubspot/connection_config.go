package hubspot

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type hubSpotConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
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
