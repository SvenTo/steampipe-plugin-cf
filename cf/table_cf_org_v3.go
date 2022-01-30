package cf

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfOrg(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_org",
		Description: "Organizations the Cloud Foundry user has access to (v3 API).",
		List: &plugin.ListConfig{
			Hydrate: listOrg,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"guid", "name"}),
			Hydrate:    getOrg,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the organization.",
				Transform:   transform.FromField("GUID"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "suspended", Type: proto.ColumnType_STRING, Description: "Whether an organization is suspended or not; non-admins will be blocked from creating, updating, or deleting resources in a suspended organization."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Labels applied and annotations added to the organization"},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links to related resources"},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationship to the quota applied to the organization"},
		},
	}
}

func listOrg(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listOrg", "connection_error", err)
		return nil, err
	}
	items, err := client.ListV3OrganizationsByQuery(url.Values{})
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listOrg", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getOrg(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getOrg", "connection_error", err)
		return nil, err
	}

	q := url.Values{}
	if name, ok := d.KeyColumnQuals["name"]; ok {
		q.Add("names", name.GetStringValue())
	} else if guid, ok := d.KeyColumnQuals["guid"]; ok {
		q.Add("guids", guid.GetStringValue())
	}

	items, err := conn.ListV3OrganizationsByQuery(q)

	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getOrg", "query_error", err)
		return nil, err
	} else if len(items) == 0 {
		return nil, nil
	}

	return items[0], err
}
