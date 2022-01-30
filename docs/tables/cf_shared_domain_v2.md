# Table: cf_shared_domain_v2

Retrieve all Cloud Foundry [shared domains](https://docs.cloudfoundry.org/devguide/routing-index.html) the user has access to ([v2](https://apidocs.cloudfoundry.org/16.22.0/domains_(deprecated)/list_all_domains_(deprecated).html)). Domains represent a fully qualified domain name that is used for application routes.

A domain can be scoped to an organization, meaning it can be used to create routes for spaces inside that organization, or be left unscoped to allow all organizations access.

## Examples

### Get all domains

```sql
select
  *
from
  cf_shared_domain_v2
```

### Get domain by guid

```sql
select
  *
from
  cf_shared_domain_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```

