# Table: cf_space_v3

Retrieve all Cloud Foundry [spaces](https://docs.cloudfoundry.org/concepts/roles.html#spaces) ([v3](http://v3-apidocs.cloudfoundry.org/version/3.113.0/index.html#spaces)) the user has access to.

## Examples

### Get all spaces the user has access to

```sql
select
  *
from
  cf_space_v3
```

### Select name and the value of the label foo

```sql
select
  name,
  metadata -> 'labels' ->> 'foo' as label_foo
from
  cf_space_v3
```

See also [Querying JSON](https://steampipe.io/docs/sql/querying-json)

### Get space by name

```sql
select
  *
from
  cf_space_v3
where
  name = 'sample_space'
```

### Get space by guid

```sql
select
  *
from
  cf_space_v3
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```