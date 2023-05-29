# Table: hubspot_hub_db

HubDB is a powerful feature in HubSpot that allows you to create and manage custom relational databases within the HubSpot platform. It provides a flexible way to store, organize, and retrieve structured data, making it suitable for a variety of use cases such as product catalogs, event listings, knowledge bases, and more.

## Examples

### Basic info

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

### List DBs created in last 30 days

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