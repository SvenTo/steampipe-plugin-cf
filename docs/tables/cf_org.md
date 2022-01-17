# Table: cf_org

Retrieve all Cloud Foundry [organizations](https://docs.cloudfoundry.org/concepts/roles.html#orgs) the user has access to.

## Examples

### Get all organization the user has access to

```sql
select
  *
from
  cf_org
```

### Get org by name

```sql
select
  *
from
  cf_org
where
  name = 'sample_org'
```

### Get org by guid

```sql
select
  *
from
  cf_org
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```