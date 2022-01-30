package cf

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfStackV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_stack_v2",
		Description: "Retrieve all stacks. Stacks are the base operating system and file system that your application will execute in (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listStackV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(250003), // cfclient error (CF-StackNotFound|250003)
			Hydrate:           getStackV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the stack",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the stack"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the stack"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was created"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was last updated"},
		},
	}
}

func listStackV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_stack_v2.listStack", "connection_error", err)
		return nil, err
	}
	items, err := client.ListStacks()
	if err != nil {
		plugin.Logger(ctx).Error("cf_stack_v2_v2.listStack", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getStackV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_stack_v2.getStack", "connection_error", err)
		return nil, err
	}

	guid := d.KeyColumnQuals["guid"].GetStringValue()

	item, err := conn.GetStackByGuid(guid)

	if err != nil {
		plugin.Logger(ctx).Error("cf_stack_v2.getStack", "query_error", err)
		return nil, err
	}

	return item, err
}
