package cf

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-cf",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		// TODO: check if correct:
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"cf_org_v2":   tableCfOrgV2(ctx),
			"cf_org_v3":   tableCfOrgV3(ctx),
			"cf_space_v3": tableCfSpaceV3(ctx),
			"cf_app_v3":   tableCfAppV3(ctx),
		},
	}
	return p
}
