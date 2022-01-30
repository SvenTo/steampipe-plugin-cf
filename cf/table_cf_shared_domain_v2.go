package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfSharedDomainV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_shared_domain_v2",
		Description: "Shared domains the user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listSharedDomainV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(130002), // cfclient error (CF-DomainNotFound|130002)
			Hydrate:           getSharedDomainV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the shared domain.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the shared domain. Must be between 3 ~ 253 characters and follow RFC 1035."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time with zone when the object was last updated."},
			{
				Name:        "router_group_guid",
				Type:        proto.ColumnType_STRING,
				Description: "The guid of the desired router group to route tcp traffic through; if set, the domain will only be available for tcp traffic.",
				Transform:   transform.FromField("RouterGroupGuid").NullIfZero(),
			},
			{Name: "router_group_type", Type: proto.ColumnType_STRING, Description: "The type of the desired router group to route tcp traffic through."},
			{Name: "internal", Type: proto.ColumnType_BOOL, Description: "Whether the domain is used for internal (container-to-container) traffic."},
		},
	}
}

func listSharedDomainV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_shared_domain_v2.listSharedDomainV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListSharedDomains()
	if err != nil {
		plugin.Logger(ctx).Error("cf_shared_domain_v2.listSharedDomainV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSharedDomainV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_shared_domain_v2.getSharedDomainV2", "connection_error", err)
		return nil, err
	}

	var item cfclient.SharedDomain
	if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetSharedDomainByGuid(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_shared_domain_v2.getSharedDomainV2", "query_error", err)
		return nil, err
	}
	return item, err
}
