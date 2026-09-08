---
title: "Steampipe Table: hubspot_team - Query HubSpot Teams using SQL"
description: "Allows users to query HubSpot Teams, providing access to the teams defined in a HubSpot account and their member users."
---

# Table: hubspot_team - Query HubSpot Teams using SQL

HubSpot Teams group users together for organization and access purposes. Each team has a set of primary member users and a set of secondary member users. Teams are referenced by users through their `primary_team_id` and `secondary_team_ids` fields.

## Table Usage Guide

The `hubspot_team` table provides insights into the teams within a HubSpot account. As a security or compliance professional, you can use this table to review team structure and membership as part of an access review. This table requires the `settings.users.teams.read` scope on the private app token.

## Examples

### Basic info
Explore the teams defined in your HubSpot account along with their member users.

```sql+postgres
select
  id,
  name,
  user_ids,
  secondary_user_ids
from
  hubspot_team;
```

```sql+sqlite
select
  id,
  name,
  user_ids,
  secondary_user_ids
from
  hubspot_team;
```

### Count primary members per team
Understand the distribution of users across teams by counting the primary members of each team.

```sql+postgres
select
  name,
  jsonb_array_length(user_ids) as member_count
from
  hubspot_team
order by
  member_count desc;
```

```sql+sqlite
select
  name,
  json_array_length(user_ids) as member_count
from
  hubspot_team
order by
  member_count desc;
```

### List teams with no members
Identify empty teams, which may be leftover structures worth cleaning up.

```sql+postgres
select
  id,
  name
from
  hubspot_team
where
  jsonb_array_length(user_ids) = 0;
```

```sql+sqlite
select
  id,
  name
from
  hubspot_team
where
  json_array_length(user_ids) = 0;
```
