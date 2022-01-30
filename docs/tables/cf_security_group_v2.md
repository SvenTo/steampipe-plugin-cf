# Table: cf_security_group_v2

Retrieve all Cloud Foundry [security groups](https://apidocs.cloudfoundry.org/16.22.0/security_groups/list_all_security_groups.html) ([v2](https://apidocs.cloudfoundry.org/16.22.0/security_groups/list_all_security_groups.html)). Security groups are collections of egress traffic rules that can be applied to the staging or running state of applications.

Security groups can either be applied globally or at the space-level.

Security groups can only allow (whitelist) traffic. They cannot be used to disallow (blacklist) traffic.

## Examples

### Get all security groups

```sql
select
  *
from
  cf_security_group_v2
```

### Get security group by guid

```sql
select
  *
from
  cf_security_group_v2
where
  guid = 'deadbeef-4242-4242-dead-beef42420001'
```

### Get security group by name

```sql
select
  *
from
  cf_security_group_v2
where
  name = 'example'
```