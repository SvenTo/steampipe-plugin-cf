# Table: cf_route_v2

Retrieve all Cloud Foundry [routes](https://docs.cloudfoundry.org/devguide/routing-index.html) the user has access to ([v2](https://apidocs.cloudfoundry.org/16.22.0/routes/list_all_routes.html)). Routes are addresses that direct matched network traffic to one or more destinations. Each route is based on a domain name with additional matching criteria (host (subdomain), path, etc).

## Examples

### Get all routes the user has access to

```sql
select
  *
from
  cf_route_v2
```

### Get route by guid

```sql
select
  *
from
  cf_route_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```
