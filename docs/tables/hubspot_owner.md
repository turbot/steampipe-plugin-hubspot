---
title: "Steampipe Table: hubspot_owner - Query HubSpot Owners using SQL"
description: "Allows users to query HubSpot Owners, providing access to detailed information about each owner in a HubSpot account."
---

# Table: hubspot_owner - Query HubSpot Owners using SQL

HubSpot Owners are the users who have been assigned ownership of specific contacts, companies, deals, or tickets in a HubSpot account. Owners can be assigned to these resources either manually or through automation, and they are responsible for managing and interacting with these resources. Each owner has a unique ID, email, and other personal details stored in the HubSpot system.

## Table Usage Guide

The `hubspot_owner` table provides insights into the owners within a HubSpot account. As a sales or marketing professional, you can use this table to explore detailed information about each owner, such as their unique ID, email, first and last name, and more. This table can be particularly useful for understanding the distribution of resources among owners in your HubSpot account.

## Examples

### Basic info
Explore the essential details of the HubSpot owners, such as their identification, title, and creation date. This is beneficial to understand the owners' archival status and personal information, which can be useful for auditing and account management purposes.

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name,
  user_id
from
  hubspot_owner;
```

### List all archived owners
Discover the segments that include all archived owners. This is useful for understanding which owners are no longer active, helping to maintain an up-to-date and accurate database.

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name,
  user_id
from
  hubspot_owner
where
  archived;
```

### List owners created in the last 30 days
Discover recent additions to your team by identifying new owners added within the past month. This is useful for keeping track of team growth and ensuring all new members are properly onboarded.

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name,
  user_id
from
  hubspot_owner
where
  created_at > now() - interval '30 days';
```

### List owners who are not associated with any team
Uncover the details of owners who are not part of any team. This can be useful for identifying potential resource allocation or communication gaps within your organization.

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name,
  user_id
from
  hubspot_owner
where
  teams is null;
```