# Table: cf_private_domain_v2

Retrieve all Cloud Foundry [private domains](https://docs.cloudfoundry.org/devguide/routing-index.html) the user has access to ([v2](https://apidocs.cloudfoundry.org/16.22.0/private_domains/list_all_private_domains.html)). Domains represent a fully qualified domain name that is used for application routes.

A domain can be scoped to an organization, meaning it can be used to create routes for spaces inside that organization, or be left unscoped to allow all organizations access.

## Examples

### Get all private domains

```sql
select
  *
from
  cf_private_domain_v2
```

### Get private domain by guid

```sql
select
  *
from
  cf_private_domain_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```
