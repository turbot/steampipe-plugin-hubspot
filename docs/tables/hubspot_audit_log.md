---
title: "Steampipe Table: hubspot_audit_log - Query HubSpot CMS Audit Logs using SQL"
description: "Allows users to query HubSpot CMS content-change audit logs, showing who created, updated, published, or deleted CMS objects."
---

# Table: hubspot_audit_log - Query HubSpot CMS Audit Logs using SQL

HubSpot CMS Audit Logs record changes made to CMS content objects such as blogs, landing pages, domains, and HubDB tables. Each entry captures the object affected, the type of change (created, updated, published, deleted, unpublished, restored), the user responsible, and when it happened.

This table covers **CMS content history only**. It is not a login or security-settings trail. For those, use `hubspot_login_activity` and `hubspot_security_activity`.

## Table Usage Guide

The `hubspot_audit_log` table provides insights into content changes within a HubSpot account. As a content administrator or auditor, you can use this table to see who changed what and when across your CMS objects. This table requires the `content` scope on the private app token. The results can be filtered by `user_id`, `object_type`, and `event`.

Note that this endpoint also requires a CMS Hub or Marketing Hub Professional or Enterprise subscription. On accounts without a qualifying tier, the HubSpot API returns a `MISSING_SCOPES` error even when the `content` scope is granted.

## Examples

### Basic info
Explore recent CMS content changes across your HubSpot account.

```sql+postgres
select
  timestamp,
  full_name,
  event,
  object_type,
  object_name
from
  hubspot_audit_log
order by
  timestamp desc
limit 10;
```

```sql+sqlite
select
  timestamp,
  full_name,
  event,
  object_type,
  object_name
from
  hubspot_audit_log
order by
  timestamp desc
limit 10;
```

### List all delete events
Focus on deletions, which are often the most sensitive content changes to review.

```sql+postgres
select
  timestamp,
  full_name,
  object_type,
  object_name
from
  hubspot_audit_log
where
  event = 'DELETED'
order by
  timestamp desc;
```

```sql+sqlite
select
  timestamp,
  full_name,
  object_type,
  object_name
from
  hubspot_audit_log
where
  event = 'DELETED'
order by
  timestamp desc;
```

### List changes made by a specific user
Review all content changes attributed to a particular user by their ID.

```sql+postgres
select
  timestamp,
  event,
  object_type,
  object_name
from
  hubspot_audit_log
where
  user_id = '12345678'
order by
  timestamp desc;
```

```sql+sqlite
select
  timestamp,
  event,
  object_type,
  object_name
from
  hubspot_audit_log
where
  user_id = '12345678'
order by
  timestamp desc;
```

### Count changes by event type
Summarize activity by the type of change made.

```sql+postgres
select
  event,
  count(*) as change_count
from
  hubspot_audit_log
group by
  event
order by
  change_count desc;
```

```sql+sqlite
select
  event,
  count(*) as change_count
from
  hubspot_audit_log
group by
  event
order by
  change_count desc;
```
