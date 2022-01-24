# Cloud Foundry Plugin for Steampipe

Use SQL to query apps, spaces and more from Cloud Foundry.

- Documentation: [Table definitions & examples](https://github.com/SvenTo/steampipe-plugin-cf/tree/main/docs/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/SvenTo/steampipe-plugin-cf/issues)

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

