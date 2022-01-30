# Table: cf_stack_v2

Retrieve all Cloud Foundry [stacks](https://docs.cloudfoundry.org/buildpacks/stack-association.html) ([v2](https://apidocs.cloudfoundry.org/16.22.0/stacks/list_all_stacks.html)). Stacks are the base operating system and file system that your application will execute in.

## Examples

### Get all stack the user has access to

```sql
select
  *
from
  cf_stack_v2
```

### Get stack by guid

```sql
select
  *
from
  cf_stack_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```

### Get stack description by app

```sql
select
  a.name as app_name,
  s.name as stack_name,
  s.description as stack_descr
from
  cf_app a,
  cf_stack_v2 s
where
  s.name = (lifecycle -> 'data' ->> 'stack')
```