---
title: "Steampipe Table: hubspot_security_activity - Query HubSpot Security Activity using SQL"
description: "Allows users to query HubSpot security activity, providing a history of security-related actions such as adding admins, enabling SSO, and installing integrations."
---

# Table: hubspot_security_activity - Query HubSpot Security Activity using SQL

HubSpot Security Activity records security-relevant actions taken in an account, such as adding or removing users, granting admin permissions, enabling single sign-on or two-factor authentication, installing integrations, and exporting CRM data. Each entry captures the type of action, the acting user, the affected object, and the source IP and location.

## Table Usage Guide

The `hubspot_security_activity` table provides insights into changes to your account's security posture and other sensitive actions. As a security professional, you can use this table to monitor privilege escalations, authentication-setting changes, and data exports. This table requires the `account-info.security.read` scope on the private app token. The results can be filtered by `user_id`.

## Examples

### Basic info
Explore recent security activity across your HubSpot account.

```sql+postgres
select
  created_at,
  type,
  acting_user,
  ip_address,
  location
from
  hubspot_security_activity
order by
  created_at desc
limit 20;
```

```sql+sqlite
select
  created_at,
  type,
  acting_user,
  ip_address,
  location
from
  hubspot_security_activity
order by
  created_at desc
limit 20;
```

### List recent admin, SSO, and two-factor authentication changes
Focus on the actions that change who holds elevated access or how users authenticate.

```sql+postgres
select
  created_at,
  type,
  acting_user,
  object_id
from
  hubspot_security_activity
where
  type in ('ADD_ADMIN_USER', 'ADD_ADMIN_PERMISSIONS', 'ADD_SINGLE_SIGN_ON', 'ADD_TWO_FACTOR_AUTHENTICATION')
order by
  created_at desc;
```

```sql+sqlite
select
  created_at,
  type,
  acting_user,
  object_id
from
  hubspot_security_activity
where
  type in ('ADD_ADMIN_USER', 'ADD_ADMIN_PERMISSIONS', 'ADD_SINGLE_SIGN_ON', 'ADD_TWO_FACTOR_AUTHENTICATION')
order by
  created_at desc;
```

### List data export activity
Review CRM export events, which move data out of HubSpot and are often subject to audit.

```sql+postgres
select
  created_at,
  acting_user,
  object_id,
  info_url
from
  hubspot_security_activity
where
  type = 'EXPORT_DOWNLOAD'
order by
  created_at desc;
```

```sql+sqlite
select
  created_at,
  acting_user,
  object_id,
  info_url
from
  hubspot_security_activity
where
  type = 'EXPORT_DOWNLOAD'
order by
  created_at desc;
```

### Count security activity by type
Summarize the volume of each kind of security action to spot unusual spikes.

```sql+postgres
select
  type,
  count(*) as activity_count
from
  hubspot_security_activity
group by
  type
order by
  activity_count desc;
```

```sql+sqlite
select
  type,
  count(*) as activity_count
from
  hubspot_security_activity
group by
  type
order by
  activity_count desc;
```
