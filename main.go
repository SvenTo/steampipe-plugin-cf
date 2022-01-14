package main

import (
	"github.com/svento/steampipe-plugin-cf/cf"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: cf.Plugin})
}
