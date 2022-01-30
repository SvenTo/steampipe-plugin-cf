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
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"cf_org_v2":  tableCfOrgV2(ctx),
			"cf_org":     tableCfOrg(ctx),
			"cf_space":   tableCfSpace(ctx),
			"cf_app":     tableCfApp(ctx),
			"cf_info_v2": tableCfInfoV2(ctx),
		},
	}
	return p
}
