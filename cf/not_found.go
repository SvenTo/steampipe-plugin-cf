package cf

import (
	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// function which returns an ErrorPredicate for Cloud Foundry API calls
func isNotFoundError(code int) plugin.ErrorPredicate {
	return func(err error) bool {
		if cfErr, ok := err.(cfclient.CloudFoundryError); ok {
			return cfErr.Code == code
		}
		return false
	}
}
