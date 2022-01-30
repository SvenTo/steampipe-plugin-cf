package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfSecGroupV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_security_group_v2",
		Description: "Security groups the Cloud Foundry user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listSecGroupV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AnyColumn([]string{"guid", "name"}),
			ShouldIgnoreError: isNotFoundError(300002), // cfclient error (CF-SecurityGroupNotFound)
			Hydrate:           getSecGroupV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the security group.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the security group."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was last updated."},
			{Name: "rules", Type: proto.ColumnType_JSON, Description: "Rules that will be applied by this security group."},
			{Name: "running", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group should be applied globally to all running applications."},
			{Name: "staging", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group should be applied globally to all staging applications."},
			{Name: "spaces_url", Type: proto.ColumnType_STRING, Description: "A URL to the spaces where the security group is applied to applications during runtime."},
			{Name: "staging_spaces_url", Type: proto.ColumnType_STRING, Description: "A URL to the spaces where the security group is applied to applications during staging."},
			{Name: "spaces_data", Type: proto.ColumnType_JSON, Description: "Data to the spaces where the security group is applied to applications during runtime."},
			{Name: "staging_spaces_data", Type: proto.ColumnType_JSON, Description: "Data to the spaces where the security group is applied to applications during staging."},
		},
	}
}

func listSecGroupV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_security_group_v2.listSecGroupV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListSecGroups()
	if err != nil {
		plugin.Logger(ctx).Error("cf_security_group_v2.listSecGroupV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSecGroupV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_security_group_v2.getSecGroupV2", "connection_error", err)
		return nil, err
	}

	var item *cfclient.SecGroup
	if name, ok := d.KeyColumnQuals["name"]; ok {
		var s cfclient.SecGroup
		s, err = conn.GetSecGroupByName(name.GetStringValue())
		item = &s
	} else if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetSecGroup(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_security_group_v2.getSecGroupV2", "query_error", err)
		return nil, err
	}
	return item, err
}
