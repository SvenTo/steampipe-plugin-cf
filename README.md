# Cloud Foundry Plugin for Steampipe

Use SQL to query apps, spaces, and more from Cloud Foundry.

- **[Get started →](https://github.com/SvenTo/steampipe-plugin-cf/blob/main/docs/index.md)**
- Documentation: [Table definitions & examples](https://github.com/SvenTo/steampipe-plugin-cf/tree/main/docs/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/SvenTo/steampipe-plugin-cf/issues)

## Quickstart

Download latest release from [GitHub](https://github.com/SvenTo/steampipe-plugin-cf/releases/).

Unzip and install the plugin:

```shell
gzip -d steampipe-plugin-cf_*.gz
chmod +x steampipe-plugin-cf_*
mkdir -p ~/.steampipe/plugins/hub.steampipe.io/plugins/svento/cf@latest/
mv steampipe-plugin-cf_* ~/.steampipe/plugins/hub.steampipe.io/plugins/svento/cf@latest/steampipe-plugin-cf.plugin
# download default config:
wget -O ~/.steampipe/config/cf.spc https://raw.githubusercontent.com/SvenTo/steampipe-plugin-cf/main/config/cf.spc
```

Log in with the [Cloud Foundry CLI](https://docs.cloudfoundry.org/cf-cli/):
```shell
cf login
```

Try it!

```
steampipe query
> .inspect cf
```

Run a query:

```sql
select
  a.name,
  a.state,
  s.name as space_name,
  o.name as org_name
from
  cf_app as a,
  cf_space as s,
  cf_org as o
where
  s.org_guid = o.guid and 
  a.space_guid = s.guid;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-cf.git
cd steampipe-plugin-cf
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/cf.spc
```

Try it!

```
cf login
steampipe query
> .inspect cf
```

Further reading about Steampipe development:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

Further reading about the Cloud Foundry API: 

- [Go reference: cloudfoundry-community/go-cfclient](https://pkg.go.dev/github.com/cloudfoundry-community/go-cfclient)
- [github.com/cloudfoundry-community/go-cfclient: Golang client lib for Cloud Foundry](https://github.com/cloudfoundry-community/go-cfclient)
- [Cloud Foundry V3 API docs](https://v3-apidocs.cloudfoundry.org/version/3.113.0/index.html#introduction)
- [Cloud Foundry V2 API docs](https://apidocs.cloudfoundry.org/16.22.0/)
- [Cloud Foundry Concepts](https://docs.cloudfoundry.org/concepts/index.html)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/SvenTo/steampipe-plugin-cf/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [CF Plugin](https://github.com/SvenTo/steampipe-plugin-cf/issues/labels/help%20wanted)
