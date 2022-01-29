package cf

import (
	"context"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCfOrgV2(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cf_org_v2",
		Description: "Organizations the Cloud Foundry user has access to (v2 API).",
		List: &plugin.ListConfig{
			Hydrate: listOrgV2,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AnyColumn([]string{"guid", "name"}),
			ShouldIgnoreError: isNotFoundError(30003), // cfclient error (CF-OrganizationNotFound|30003)
			Hydrate:           getOrgV2,
		},
		Columns: []*plugin.Column{
			{
				Name:        "guid",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the organization.",
				Transform:   transform.FromField("Guid"),
			},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the organization."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was created."},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: "The time with zone when the object was last updated."},
			{
				Name:        "quota_definition_guid",
				Type:        proto.ColumnType_STRING,
				Description: "The guid of quota to associate with this organization.",
				Transform:   transform.FromField("QuotaDefinitionGuid"),
			},
			{
				Name:        "default_isolation_segment_guid",
				Type:        proto.ColumnType_STRING,
				Description: "The guid of the isolation segment to set as the organizational default.",
				Transform:   transform.FromField("DefaultIsolationSegmentGuid"),
			},
		},
	}
}

func listOrgV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org_v2.listOrgV2", "connection_error", err)
		return nil, err
	}
	items, err := client.ListOrgs()
	if err != nil {
		plugin.Logger(ctx).Error("cf_org_v2.listOrgV2", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getOrgV2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cf_org_v2.getOrgV2", "connection_error", err)
		return nil, err
	}

	var item cfclient.Org
	if name, ok := d.KeyColumnQuals["name"]; ok {
		item, err = conn.GetOrgByName(name.GetStringValue())
	} else if guid, ok := d.KeyColumnQuals["guid"]; ok {
		item, err = conn.GetOrgByGuid(guid.GetStringValue())
	}

	if err != nil {
		plugin.Logger(ctx).Error("cf_org_v2.getOrgV2", "query_error", err)
		return nil, err
	}
	return item, err
}
