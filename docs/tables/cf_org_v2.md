# Table: cf_org_v2

Retrieve all Cloud Foundry [organizations](https://docs.cloudfoundry.org/concepts/roles.html#orgs) the user has access to (v2 API).

## Examples

### Get all organization the user has access to

```sql
select
  *
from
  cf_org_v2
```

### Get org by name

```sql
select
  *
from
  cf_org_v2
where
  name = 'sample_org'
```

### Get org by guid

```sql
select
  *
from
  cf_org_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```