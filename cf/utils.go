package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*cfclient.Client, error) {

	// Load connection from cache, to
	// TODO: do not do this if config is loaded from CF home configuration file
	cacheKey := "cf"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*cfclient.Client), nil
	}


	c, err := cfclient.NewConfigFromCF()

	if err != nil {
		return nil, err
	}

	conn, err := cfclient.NewClient(c)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
