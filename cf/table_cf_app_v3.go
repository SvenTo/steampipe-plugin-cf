package cf

import (
	"context"
	"net/url"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfAppV3(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_app_v3",
		Description: "Apps the user has access to (v3 API).",
		List: &plugin.ListConfig{
			Hydrate: listAppV3,
		},
		Get: &plugin.GetConfig{
			// TODO: Add organization_guid
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(30003), // cfclient error (CF-OrganizationNotFound|30003)
			Hydrate:           getAppV3,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the app",
				Transform:   transform.FromField("GUID"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the app"},
			{
				Name:        "space_guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the space.",
				Transform:   transform.From(transformSpaceGuid),
			},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created"},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated"},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "Current desired state of the app; valid values are STOPPED or STARTED"},
			{Name: "lifecycle", Type: proto.ColumnType_JSON, Description: "Provides the default lifecycle object for the application. This lifecycle will be used when staging and running the application. The staging lifecycle can be overridden on builds."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Labels applied and annotations added to the app"},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links to related resources"},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationship to the space the app is contained in"},
		},
	}
}

func listAppV3(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listAppV3", "connection_error", err)
		return nil, err
	}
	items, err := client.ListV3AppsByQuery(url.Values{})
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.listAppV3", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getAppV3(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getAppV3", "connection_error", err)
		return nil, err
	}

	q := url.Values{}
	q.Add("guids", d.KeyColumnQuals["guid"].GetStringValue())

	items, err := conn.ListV3AppsByQuery(q)

	if err != nil {
		plugin.Logger(ctx).Error("cf_org.getAppV3", "query_error", err)
		return nil, err
	}
	// TODO: check for length
	return items[0], err
}

//// TRANSFORM FUNCTION

func transformSpaceGuid(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(cfclient.V3App)
	return data.Relationships["space"].Data.GUID, nil
}