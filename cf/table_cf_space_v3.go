package cf

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfSpaceV3(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_space_v3",
		Description: "Spaces the Cloud Foundry user has access to (v3 API).",
		List: &plugin.ListConfig{
			Hydrate: listSpaceV3,
		},
		Get: &plugin.GetConfig{
			// TODO: Add organization_guid
			KeyColumns:        plugin.AnyColumn([]string{"guid", "name"}),
			ShouldIgnoreError: isNotFoundError(30003), // cfclient error (CF-OrganizationNotFound|30003)
			Hydrate:           getSpaceV3,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the space",
				Transform:   transform.FromField("GUID"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the space"},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created"},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated"},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Labels applied and annotations added to the space"},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links to related resources"},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationship to the quota applied to the space and the organization the space is contained in"},
			// TODO: Add organization_guid
		},
	}
}

func listSpaceV3(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listSpaceV3", "connection_error", err)
		return nil, err
	}
	items, err := client.ListV3SpacesByQuery(url.Values{})
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listSpaceV3", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSpaceV3(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getSpaceV3", "connection_error", err)
		return nil, err
	}

	q := url.Values{}
	if name, ok := d.KeyColumnQuals["name"]; ok {
		q.Add("names", name.GetStringValue())
	} else if guid, ok := d.KeyColumnQuals["guid"]; ok {
		q.Add("guids", guid.GetStringValue())
	}
	// TODO: Add organization_guid

	items, err := conn.ListV3SpacesByQuery(q)

	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getSpaceV3", "query_error", err)
		return nil, err
	}
	// TODO: Check length for >0
	return items[0], err
}
