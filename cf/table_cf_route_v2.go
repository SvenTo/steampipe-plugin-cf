package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfRouteV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_route_v2",
		Description: "Routes the Cloud Foundry user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listRouteV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(210002), // cfclient error (CF-RouteNotFound|210002)
			Hydrate:           getRouteV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the route.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "host", Type: proto.ColumnType_STRING, Description: "The hostname for the route; not compatible with routes specifying the tcp protocol;"},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "The path for the route; not compatible with routes specifying the tcp protocol;", Transform: transform.FromGo()},
			{
				Name: "domain_guid", Type: proto.ColumnType_STRING,
				Description: "The guid to the domain of the route.",
				Transform:   transform.FromField("DomainGuid"),
			},
			{Name: "domain_url", Type: proto.ColumnType_STRING, Description: "Status of the route."},
			{
				Name: "space_guid", Type: proto.ColumnType_STRING,
				Description: "The guid to the space of the route.",
				Transform:   transform.FromField("SpaceGuid"),
			},
			{
				Name: "service_instance_guid", Type: proto.ColumnType_STRING,
				Description: "The guid to the service instance of the route.",
				Transform:   transform.FromField("ServiceInstanceGuid").NullIfZero(),
			},
			{Name: "port", Type: proto.ColumnType_INT, Description: "The port that the route listens on. Only compatible with routes specifying the tcp protocol."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was last updated."},
		},
	}
}

func listRouteV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_route_v2.listRouteV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListRoutes()
	if err != nil {
		plugin.Logger(ctx).Error("cf_route_v2.listRouteV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getRouteV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_route_v2.getRouteV2", "connection_error", err)
		return nil, err
	}

	var item cfclient.Route
	if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetRouteByGuid(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_route_v2.getRouteV2", "query_error", err)
		return nil, err
	}
	return item, err
}
