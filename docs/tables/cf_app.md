# Table: cf_app

Retrieve all Cloud Foundry [apps](https://docs.cloudfoundry.org/devguide/deploy-apps/) ([v3](http://v3-apidocs.cloudfoundry.org/version/3.113.0/index.html#apps)) the user has access to.

## Examples

### Get all apps the user has access to

```sql
select
  *
from
  cf_app
```

### Get all apps by an specific space

```sql
select
  a.*
from
  cf_app as a,
  cf_space as s
where
  s.name = 'dev' and
  a.space_guid = s.guid
```

### Select name and the value of the label foo

```sql
select
  name,
  metadata -> 'labels' ->> 'foo' as label_foo
from
  cf_app
```

See also [Querying JSON](https://steampipe.io/docs/sql/querying-json)

### Get apps by name

```sql
select
  *
from
  cf_app
where
  name = 'sample_app'
```

### Get app by guid

```sql
select
  *
from
  cf_app
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```