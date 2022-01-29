package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*cfclient.Client, error) {
	var err error

	// Try to load connection from cache
	cacheKey := "cf"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*cfclient.Client), nil
	}

	spConfig := GetConfig(d.Connection)
	c, cacheConfig, err := spConfig.convertToCFClientConfig()

	if err != nil {
		return nil, err
	}

	conn, err := cfclient.NewClient(c)

	if err != nil {
		return nil, err
	}

	// In default we don't cache the config so that the most recent auth token is loaded from the CF config file (~/.cf/config.json)
	if cacheConfig {
		// Save to cache
		d.ConnectionManager.Cache.Set(cacheKey, conn)
	}

	return conn, nil
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
