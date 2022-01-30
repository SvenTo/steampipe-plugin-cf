package cf

import (
	"context"
	"net/url"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfSpace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_space",
		Description: "Spaces the Cloud Foundry user has access to (v3 API).",
		List: &plugin.ListConfig{
			Hydrate: listSpace,
		},
		Get: &plugin.GetConfig{
			// name cannot be a key column because it is not unique across orgs
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(30003), // cfclient error (CF-OrganizationNotFound|30003)
			Hydrate:           getSpace,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the space",
				Transform:   transform.FromField("GUID"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the space"},
			{
				Name:        "org_guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the organization.",
				Transform:   transform.From(transformOrganizationGuid),
			},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created"},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated"},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Labels applied and annotations added to the space"},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links to related resources"},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationship to the quota applied to the space and the organization the space is contained in"},
		},
	}
}

func listSpace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_space.listSpace", "connection_error", err)
		return nil, err
	}

	items, err := client.ListV3SpacesByQuery(url.Values{})
	if err != nil {
		plugin.Logger(ctx).Error("cf_space.listSpace", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSpace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_space.getSpace", "connection_error", err)
		return nil, err
	}

	q := url.Values{}
	q.Add("guids", d.KeyColumnQuals["guid"].GetStringValue())

	items, err := conn.ListV3SpacesByQuery(q)

	if err != nil {
		plugin.Logger(ctx).Error("cf_space.getSpace", "query_error", err)
		return nil, err
	} else if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

//// TRANSFORM FUNCTION

func transformOrganizationGuid(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(cfclient.V3Space)
	return data.Relationships["organization"].Data.GUID, nil
}
