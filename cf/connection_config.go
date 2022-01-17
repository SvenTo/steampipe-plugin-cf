package cf

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type cfConfig struct {
	// TODO: replace with actual config parameters
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	// TODO: replace with actual config parameters
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &cfConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) cfConfig {
	if connection == nil || connection.Config == nil {
		return cfConfig{}
	}
	config, _ := connection.Config.(cfConfig)
	return config
}
