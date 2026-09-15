---
title: "Steampipe Table: hubspot_user - Query HubSpot Users using SQL"
description: "Allows users to query HubSpot Users, providing access to their roles, teams, and super admin status for access reviews."
---

# Table: hubspot_user - Query HubSpot Users using SQL

HubSpot Users are the people who can sign in to a HubSpot account. Each user has an email address, an assigned role, primary and secondary team memberships, and may be flagged as a super admin. This information is the foundation for reviewing who has access to an account and at what privilege level.

## Table Usage Guide

The `hubspot_user` table provides insights into the users within a HubSpot account. As a security or compliance professional, you can use this table to review account membership, identify super admins, and map users to their roles and teams. This table requires the `settings.users.read` scope on the private app token.

## Examples

### Basic info
Explore the users in your HubSpot account along with their role and super admin status. This helps you understand who has access to the account and at what level.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  super_admin,
  role_id
from
  hubspot_user;
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  super_admin,
  role_id
from
  hubspot_user;
```

### List all super admins
Identify the users with super admin privileges. Super admins have unrestricted access, so keeping this list small and known is a common security control.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name
from
  hubspot_user
where
  super_admin;
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name
from
  hubspot_user
where
  super_admin = 1;
```

### List users not assigned to any team
Uncover users who do not belong to a primary team. This can highlight gaps in team-based access organization.

```sql+postgres
select
  id,
  email,
  role_id
from
  hubspot_user
where
  primary_team_id is null;
```

```sql+sqlite
select
  id,
  email,
  role_id
from
  hubspot_user
where
  primary_team_id is null;
```

### Count users per role
Understand how privileges are distributed across your account by counting users per assigned role.

```sql+postgres
select
  role_id,
  count(*) as user_count
from
  hubspot_user
group by
  role_id
order by
  user_count desc;
```

```sql+sqlite
select
  role_id,
  count(*) as user_count
from
  hubspot_user
group by
  role_id
order by
  user_count desc;
```

### Get details for a specific user
Retrieve the full record for a single user by their ID.

```sql+postgres
select
  id,
  email,
  super_admin,
  role_ids,
  secondary_team_ids
from
  hubspot_user
where
  id = '12345678';
```

```sql+sqlite
select
  id,
  email,
  super_admin,
  role_ids,
  secondary_team_ids
from
  hubspot_user
where
  id = '12345678';
```
