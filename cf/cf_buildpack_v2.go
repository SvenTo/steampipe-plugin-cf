package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfBuildpackV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_buildpack_v2",
		Description: "Retrieve all builds the Cloud Foundry user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listBuildpackV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AnyColumn([]string{"guid", "name"}),
			ShouldIgnoreError: isNotFoundError(10000), // cfclient error (CF-NotFound|10000)
			Hydrate:           getBuildpackV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the buildpack.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the buildpack; to be used by app buildpack field (only alphanumeric characters)"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The state of the buildpack."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was last updated."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "Whether or not the buildpack can be used for staging."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "Whether or not the buildpack is locked to prevent updating the bits."},
			{Name: "position", Type: proto.ColumnType_INT, Description: "The order in which the buildpacks are checked during buildpack auto-detection."},
			{Name: "filename", Type: proto.ColumnType_STRING, Description: "The filename of the buildpack."},
			{Name: "stack", Type: proto.ColumnType_STRING, Description: "The name of the stack that the buildpack will use."},
		},
	}
}

func listBuildpackV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_buildpack_v2.listBuildpackV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListBuildpacks()
	if err != nil {
		plugin.Logger(ctx).Error("cf_buildpack_v2.listBuildpackV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getBuildpackV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_buildpack_v2.getBuildpackV2", "connection_error", err)
		return nil, err
	}

	var item cfclient.Buildpack
	if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetBuildpackByGuid(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_buildpack_v2.getBuildpackV2", "query_error", err)
		return nil, err
	}
	return item, err
}
