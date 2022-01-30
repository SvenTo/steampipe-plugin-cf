package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfPrivateDomainV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_private_domain_v2",
		Description: "Private domains the user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listPrivateDomainV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("guid"),
			ShouldIgnoreError: isNotFoundError(130002), // cfclient error (CF-DomainNotFound|130002)
			Hydrate:           getPrivateDomainV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the private domain.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the private domain. Must be between 3 ~ 253 characters and follow RFC 1035."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated."},
			{
				Name:        "owning_organization_guid",
				Type:        proto.ColumnType_STRING,
				Description: "The guid of the owning organization the domain is scoped to.",
				Transform:   transform.FromField("OwningOrganizationGuid"),
			},
			{Name: "owning_organization_url", Type: proto.ColumnType_STRING, Description: "The organization the domain is scoped to; if set, the domain will only be available in that organization; otherwise, the domain will be globally available."},
			{Name: "shared_organizations_url", Type: proto.ColumnType_STRING, Description: "Organizations the domain is shared with; if set, the domain will be available in these organizations in addition to the organization the domain is scoped to."},
		},
	}
}

func listPrivateDomainV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_private_domain_v2.listPrivateDomainV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListDomains()
	if err != nil {
		plugin.Logger(ctx).Error("cf_private_domain_v2.listPrivateDomainV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getPrivateDomainV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_private_domain_v2.getPrivateDomainV2", "connection_error", err)
		return nil, err
	}

	var item cfclient.Domain
	if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetDomainByGuid(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_private_domain_v2.getPrivateDomainV2", "query_error", err)
		return nil, err
	}
	return item, err
}
