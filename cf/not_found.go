package cf

import (
	"errors"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// function which returns an ErrorPredicate for Cloud Foundry API calls
func isNotFoundError(code int) plugin.ErrorPredicate {
	return func(err error) bool {
		var cfErr cfclient.CloudFoundryError
		// some CloudFoundryError errors are wrapped in another error:
		ok := errors.As(err, &cfErr)
		if ok {
			return cfErr.Code == code
		}
		return false
	}
}
