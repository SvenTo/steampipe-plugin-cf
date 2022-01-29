---
organization: SvenTo
category: ["software development"]
icon_url: "/images/plugins/svento/cf.svg"
brand_color: "#a7cae1"
display_name: "Cloud Foundry"
short_name: "cf"
description: "Steampipe plugin to query apps, spaces and more from Cloud Foundry."
og_description: "Query Cloud Foundry with SQL! Open source CLI. No DB required."
---

# Cloud Foundry + Steampipe

[Cloud Foundry](https://www.cloudfoundry.org/)  is an open source, multi-cloud application platform as a service (PaaS).

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List apps in your Cloud Foundry account:

```sql
select
  a.name,
  a.state,
  s.name as space_name,
  o.name as org_name
from
  cf_app_v3 as a,
  cf_space_v3 as s,
  cf_org_v3 as o
where
  s.org_guid = o.guid and 
  a.space_guid = s.guid;
```

```
+--------------+---------+------------+---------------+
| name         | state   | space_name | org_name      |
+--------------+---------+------------+---------------+
| spring-music | STARTED | dev2       | sample_org    |
| spring-store | STOPPED | dev        | sample_org    |
+--------------+---------+------------+---------------+
```

## Documentation

- **[Table definitions & examples →](https://github.com/SvenTo/steampipe-plugin-cf/tree/main/docs/tables)**

## Get started

### Configuration

Installing the latest cf plugin will create a config file (`~/.steampipe/config/cf.spc`) with a single connection named `cf`:

```hcl
connection "cf" {
  plugin      = "cf"
}
```

This will create a `cf` connection that uses the [Cloud Foundry CLI configuration](https://docs.cloudfoundry.org/cf-cli/) file from default config directory (typically ``~/.cf/config.json`` on *nix systems). Depending on your configuration, you may need to be [logged in](https://docs.cloudfoundry.org/cf-cli/getting-started.html#login) with the CF CLI. Use [``cf login``](https://docs.cloudfoundry.org/cf-cli/getting-started.html#login) if you get the error ``(CF-InvalidAuthToken|1000): Invalid Auth Token``.


## Get involved

- Open source: https://github.com/svento/steampipe-plugin-cf
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
