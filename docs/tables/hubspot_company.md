# Table: hubspot_company

In HubSpot, the Company object represents a company or organization with which you interact or have a business relationship. It is a core entity in HubSpot's CRM (Customer Relationship Management) system and serves as a central repository for storing and managing information related to your company contacts, deals, interactions, and more.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company;
```

### List all archived companies

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  archived;
```

### List publicly traded companies

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  is_public;
```

### List companies located in Delhi

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  state = 'Delhi';
```

### List vendor companies

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  type = 'VENDOR';
```

### List companies with less than 200 employees

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  numberofemployees < '200';
```

### List IT Services companies

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  industry = 'INFORMATION_TECHNOLOGY_AND_SERVICES';
```

### List companies created in the last 30 days

```sql
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  created_at > now() - interval '30 days';
```
