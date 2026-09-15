---
title: "Steampipe Table: hubspot_login_activity - Query HubSpot Login Activity using SQL"
description: "Allows users to query HubSpot login activity, providing a history of successful and failed sign-in attempts over the past 90 days."
---

# Table: hubspot_login_activity - Query HubSpot Login Activity using SQL

HubSpot Login Activity records user sign-in attempts, both successful and unsuccessful, made in the past 90 days. This includes logins to app.hubspot.com and the HubSpot mobile app. Each entry captures the user, whether the attempt succeeded, the source IP address, and the approximate location.

## Table Usage Guide

The `hubspot_login_activity` table provides insights into who has been signing in to your HubSpot account and from where. As a security professional, you can use this table to detect failed login attempts, logins from unexpected locations, and other anomalous sign-in behavior. This table requires the `account-info.security.read` scope on the private app token. The results can be filtered by `user_id`.

## Examples

### Basic info
Explore recent login activity across your HubSpot account.

```sql+postgres
select
  login_at,
  email,
  login_succeeded,
  ip_address,
  location
from
  hubspot_login_activity
order by
  login_at desc
limit 20;
```

```sql+sqlite
select
  login_at,
  email,
  login_succeeded,
  ip_address,
  location
from
  hubspot_login_activity
order by
  login_at desc
limit 20;
```

### List failed login attempts
Focus on unsuccessful sign-in attempts, which can indicate credential-guessing or account-takeover attempts.

```sql+postgres
select
  login_at,
  email,
  ip_address,
  location
from
  hubspot_login_activity
where
  not login_succeeded
order by
  login_at desc;
```

```sql+sqlite
select
  login_at,
  email,
  ip_address,
  location
from
  hubspot_login_activity
where
  login_succeeded = 0
order by
  login_at desc;
```

### List logins from outside a specific country
Identify sign-ins that originate from unexpected countries as part of a geo-anomaly review.

```sql+postgres
select
  login_at,
  email,
  ip_address,
  location,
  country_code
from
  hubspot_login_activity
where
  country_code <> 'us'
order by
  login_at desc;
```

```sql+sqlite
select
  login_at,
  email,
  ip_address,
  location,
  country_code
from
  hubspot_login_activity
where
  country_code <> 'us'
order by
  login_at desc;
```

### Count failed logins per user
Summarize failed attempts by user to spot accounts under repeated pressure.

```sql+postgres
select
  email,
  count(*) as failed_attempts
from
  hubspot_login_activity
where
  not login_succeeded
group by
  email
order by
  failed_attempts desc;
```

```sql+sqlite
select
  email,
  count(*) as failed_attempts
from
  hubspot_login_activity
where
  login_succeeded = 0
group by
  email
order by
  failed_attempts desc;
```
