# Table: cf_buildpack_v2

Retrieve all Cloud Foundry [buildpacks](https://docs.cloudfoundry.org/buildpacks/stack-association.html) ([v2](https://apidocs.cloudfoundry.org/16.22.0/stacks/list_all_stacks.html)) the user has access to. Buildpacks are used during a build to download external dependencies and transform a package into an executable droplet. In this way, buildpacks are a pluggable extension to Cloud Foundry that enable CF to run different languages and frameworks. Buildpacks will automatically detect if they support an application. Buildpacks can also be explicitly specified on apps and builds.

## Examples

### Get all buildpacks

```sql
select
  *
from
  cf_buildpack_v2
```

### Get buildpack by guid

```sql
select
  *
from
  cf_buildpack_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```
