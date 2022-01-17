# Table: cf_org_v3

Retrieve all Cloud Foundry [organizations](https://docs.cloudfoundry.org/concepts/roles.html#orgs) ([v3](http://v3-apidocs.cloudfoundry.org/version/3.113.0/index.html#organizations)) the user has access to.

## Examples

### Get all organization the user has access to

```sql
select
  *
from
  cf_org_v3
```

### Select name and the value of the label foo

```sql
select
  name,
  metadata -> 'labels' ->> 'foo' as label_foo
from
  cf_org_v3
```

See also [Querying JSON](https://steampipe.io/docs/sql/querying-json)

### Get org by name

```sql
select
  *
from
  cf_org_v3
where
  name = 'sample_org'
```

### Get org by guid

```sql
select
  *
from
  cf_org_v3
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```