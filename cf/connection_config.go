package cf

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type cfConfig struct {
	// TODO: replace with actual config parameters
	CFHomeDir         *string `cty:"cf_home_dir"`
	ApiAddress        *string `cty:"api_url"`
	Username          *string `cty:"user"`
	Password          *string `cty:"password"`
	ClientID          *string `cty:"client_id"`
	ClientSecret      *string `cty:"client_secret"`
	Token             *string `cty:"auth_token"`
	SkipSslValidation *bool   `cty:"skip_ssl_validation"`
	UserAgent         *string `cty:"user_agent"`
}

var ConfigSchema = map[string]*schema.Attribute{
	// TODO: replace with actual config parameters
	"cf_home_dir": {
		Type: schema.TypeString,
	},
	"api_url": {
		Type: schema.TypeString,
	},
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
	"skip_ssl_validation": {
		Type: schema.TypeBool,
	},
	"auth_token": {
		Type: schema.TypeString,
	},
	"user_agent": {
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

func (sp cfConfig) convertToCFClientConfig() (*cfclient.Config, bool, error) {
	var c *cfclient.Config
	var err error
	cacheConfig := false

	if sp.ApiAddress != nil {
		c = cfclient.DefaultConfig()
		c.ApiAddress = derefString(sp.ApiAddress)
		c.Username = derefString(sp.Username)
		c.Password = derefString(sp.Password)
		c.ClientID = derefString(sp.ClientID)
		c.ClientSecret = derefString(sp.ClientSecret)
		c.Token = derefString(sp.Token)
		cacheConfig = true
	} else {
		c, err = cfclient.NewConfigFromCFHome(derefString(sp.CFHomeDir))
	}

	if err != nil {
		return nil, cacheConfig, err
	}

	if sp.UserAgent != nil {
		c.UserAgent = *sp.UserAgent
	}
	if sp.SkipSslValidation != nil {
		c.SkipSslValidation = *sp.SkipSslValidation
	}

	return c, cacheConfig, nil
}
