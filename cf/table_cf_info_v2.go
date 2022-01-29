package cf

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableCfInfoV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_info_v2",
		Description: "Get Cloud Foundry platform info.",
		List: &plugin.ListConfig{
			Hydrate: listInfoV2,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the platform"},
			{Name: "build", Type: proto.ColumnType_STRING, Description: "Build version of the platform"},
			{Name: "support", Type: proto.ColumnType_STRING, Description: "Support URL"},
			{Name: "version", Type: proto.ColumnType_INT, Description: "Version of the platform"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the platform"},
			{Name: "authorization_endpoint", Type: proto.ColumnType_STRING, Description: "URL of the authorization endpoint"},
			{Name: "token_endpoint", Type: proto.ColumnType_STRING, Description: "URL of the token endpoint (UAA)"},
			{Name: "min_cli_version", Type: proto.ColumnType_STRING, Description: "Minimum required CF CLI version"},
			{Name: "min_recommended_cli_version", Type: proto.ColumnType_STRING, Description: "Minimum recommended CF CLI version"},
			{Name: "api_version", Type: proto.ColumnType_STRING, Description: "API version"},
			{Name: "app_ssh_endpoint", Type: proto.ColumnType_STRING, Description: "SSH endpoint in hostname:port format"},
			{Name: "app_ssh_host_key_fingerprint", Type: proto.ColumnType_STRING, Description: "SSH host key fingerprint"},
			{Name: "app_ssh_oauth_client", Type: proto.ColumnType_STRING, Description: "SSH OAuth client"},
			{Name: "doppler_logging_endpoint", Type: proto.ColumnType_STRING, Description: "Doppler logging endpoint"},
			{Name: "routing_endpoint", Type: proto.ColumnType_STRING, Description: "Routing endpoint"},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "User ID or null if unauthenticated"},
		},
	}
}

func listInfoV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_info_v2.listInfo", "connection_error", err)
		return nil, err
	}

	item, err := client.GetInfo()
	if err != nil {
		plugin.Logger(ctx).Error("cf_info_v2.listInfo", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, item)

	return nil, nil
}
