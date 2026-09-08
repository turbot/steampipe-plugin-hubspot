---
title: "Steampipe Table: hubspot_user_role - Query HubSpot User Roles using SQL"
description: "Allows users to query HubSpot User Roles, providing access to the roles defined in a HubSpot account."
---

# Table: hubspot_user_role - Query HubSpot User Roles using SQL

HubSpot User Roles are the named permission sets that can be assigned to users in a HubSpot account. Roles are referenced by users through their `role_id` and `role_ids` fields. Note that the HubSpot API exposes only the role names and IDs, not the granular permission grants each role contains.

## Table Usage Guide

The `hubspot_user_role` table provides insights into the roles defined within a HubSpot account. As a security or compliance professional, you can use this table to enumerate the roles available for assignment and to join against `hubspot_user` to see which users hold which role. This table requires the `settings.users.read` scope on the private app token.

Note that the roles feature requires a qualifying HubSpot subscription. On accounts without access to roles, the HubSpot API returns an error (`Account doesn't have access to roles`) rather than an empty list.

## Examples

### Basic info
Explore the roles defined in your HubSpot account.

```sql+postgres
select
  id,
  name,
  requires_billing_write
from
  hubspot_user_role;
```

```sql+sqlite
select
  id,
  name,
  requires_billing_write
from
  hubspot_user_role;
```

### List roles that require billing write access
Identify the roles that grant billing write access, which is a sensitive permission worth reviewing.

```sql+postgres
select
  id,
  name
from
  hubspot_user_role
where
  requires_billing_write;
```

```sql+sqlite
select
  id,
  name
from
  hubspot_user_role
where
  requires_billing_write = 1;
```

### List users along with their role names
Join users to roles to produce a readable access review that maps each user to their assigned role name.

```sql+postgres
select
  u.email,
  u.super_admin,
  r.name as role_name
from
  hubspot_user as u
  left join hubspot_user_role as r on u.role_id = r.id
order by
  r.name;
```

```sql+sqlite
select
  u.email,
  u.super_admin,
  r.name as role_name
from
  hubspot_user as u
  left join hubspot_user_role as r on u.role_id = r.id
order by
  r.name;
```
