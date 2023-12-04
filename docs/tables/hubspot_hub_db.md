---
title: "Steampipe Table: hubspot_hub_db - Query HubSpot Hubs using SQL"
description: "Allows users to query HubSpot Hubs, providing comprehensive details about the hubs including their ID, name, portal ID, and other relevant information."
---

# Table: hubspot_hub_db - Query HubSpot Hubs using SQL

HubSpot Hubs are a part of the HubSpot platform designed to help businesses grow by bringing together different functionalities such as marketing, sales, service, and CRM. Each Hub is a suite of related tools, and they can be used individually or combined for a fully integrated experience. The HubSpot platform helps businesses attract visitors, convert leads, close customers, and delight customers into promoters.

HubDB is a powerful feature in HubSpot that allows you to create and manage custom relational databases within the HubSpot platform. It provides a flexible way to store, organize, and retrieve structured data, making it suitable for a variety of use cases such as product catalogs, event listings, knowledge bases, and more.

## Table Usage Guide

The `hubspot_hub_db` table provides insights into HubSpot Hubs within the HubSpot platform. As a marketing or sales professional, explore hub-specific details through this table, including the ID, name, and portal ID. Utilize it to uncover comprehensive information about each hub, such as its associated portal, helping you to better manage and organize your marketing and sales efforts.

## Examples

### Basic info
Explore the basic information about your Hubspot database, such as the number of columns and rows, publishing status, and creation details. This can be useful to understand the overall structure and content of your database at a glance.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db;
```

### List archived DBs
Explore which databases have been archived. This can be useful for maintaining data integrity and ensuring that outdated or unnecessary information is properly stored away.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db
where
  archived;
```

### List public DBs
Explore which databases are set to public access to understand potential security risks and ensure appropriate data privacy measures are in place. This can be particularly useful for auditing and compliance purposes.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db
where
  allow_public_api_access;
```

### List published DBs
Explore which databases have been published in your Hubspot Hub. This can help you understand the scope and scale of your data, as well as track the creation and publication of new databases.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db
where
  published;
```

### List DBs where child tables are not allowed
Discover the segments where child tables are disallowed in databases. This can be useful to understand the structure and restrictions of your database system.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db
where
  not allow_child_tables;
```

### List DBs created in the last 30 days
Discover the databases that were established within the past month. This query is beneficial for maintaining a fresh perspective on recent data additions and understanding the growth rate of your database accumulation.

```sql
select
  id,
  title,
  label,
  published,
  column_count,
  row_count,
  created_by,
  created_at
from
  hubspot_hub_db
where
  created_at > now() - interval '30 days';
```
