---
title: "Steampipe Table: hubspot_access_token_info - Query HubSpot Access Token Info using SQL"
description: "Allows users to query the private app access token in use, including the scopes it has been granted."
---

# Table: hubspot_access_token_info - Query HubSpot Access Token Info using SQL

HubSpot Access Token Info describes the private app access token that the connection is currently using. It introspects the token to reveal the account (hub) and app it belongs to, the user it is associated with, and, most importantly, the scopes it has been granted. This is useful for verifying that a token follows least-privilege principles.

## Table Usage Guide

The `hubspot_access_token_info` table provides insights into the credential the plugin is using. As a security or compliance professional, you can use this table to audit the scopes granted to the token and confirm it is not over-provisioned. This table introspects the token itself and requires no additional scope. It returns a single row.

## Examples

### Basic info
Explore the account, app, and user associated with the token in use.

```sql+postgres
select
  hub_id,
  app_id,
  user_id,
  is_user_token
from
  hubspot_access_token_info;
```

```sql+sqlite
select
  hub_id,
  app_id,
  user_id,
  is_user_token
from
  hubspot_access_token_info;
```

### List the scopes granted to the token
Review the full set of scopes the token holds. Use this to confirm the token has only the access it needs.

```sql+postgres
select
  scope
from
  hubspot_access_token_info,
  jsonb_array_elements_text(scopes) as scope
order by
  scope;
```

```sql+sqlite
select
  scope.value as scope
from
  hubspot_access_token_info,
  json_each(scopes) as scope
order by
  scope.value;
```

### Check whether the token can read users
Confirm whether the token holds a specific sensitive scope, such as the ability to read account users.

```sql+postgres
select
  hub_id,
  scopes ? 'settings.users.read' as can_read_users
from
  hubspot_access_token_info;
```

```sql+sqlite
select
  hub_id,
  exists (
    select 1
    from json_each(scopes) as scope
    where scope.value = 'settings.users.read'
  ) as can_read_users
from
  hubspot_access_token_info;
```
