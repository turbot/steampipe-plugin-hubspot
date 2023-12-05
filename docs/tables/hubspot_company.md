---
title: "Steampipe Table: hubspot_company - Query HubSpot Companies using SQL"
description: "Allows users to query HubSpot Companies, providing detailed information on various aspects of a company such as its name, website, number of employees, and more."
---

# Table: hubspot_company - Query HubSpot Companies using SQL

HubSpot Companies is a resource within HubSpot that represents a collection of individuals who you do business with. Companies in HubSpot are distinct entities, separate from contacts, but they can be associated with multiple contacts. This allows for easy tracking and organization of business interactions.

## Table Usage Guide

The `hubspot_company` table provides insights into companies within HubSpot. As a sales or marketing professional, explore company-specific details through this table, including company size, industry, and associated contacts. Utilize it to uncover information about companies, such as their primary business details, revenue data, and the relationships with contacts.

## Examples

### Basic info
Explore the creation and modification dates of companies in your Hubspot account, including whether they've been archived, to assess changes over time and manage your contacts effectively.

```sql+postgres
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company;
```

```sql+sqlite
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
Discover the segments that consist of archived companies, allowing you to analyze and understand the historical data associated with these businesses. This can be particularly helpful in assessing business trends and patterns over time.

```sql+postgres
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

```sql+sqlite
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  archived = 1;
```

### List publicly traded companies
Explore which companies in your Hubspot database are publicly traded. This can be beneficial for understanding your customer base and identifying potential investment opportunities.

```sql+postgres
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

```sql+sqlite
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  is_public = 1;
```

### List companies located in Delhi
Explore which businesses are based in Delhi. This is particularly useful for market research, competitor analysis, or potential partnership opportunities.

```sql+postgres
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

```sql+sqlite
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
Explore which companies are classified as vendors within your Hubspot database. This is useful for gaining insights into your vendor relationships and assessing their status.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that consist of companies with less than 200 employees. This can be useful in assessing the size and scale of potential business partners or competitors.

```sql+postgres
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

```sql+sqlite
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  numberofemployees < 200;
```

### List IT Services companies
Explore which companies in the Information Technology and Services industry are available. This is useful for identifying potential business partners or competitors in the same industry.

```sql+postgres
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

```sql+sqlite
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
Explore which companies have been established within the past month. This can be useful to track new business growth and potential opportunities.

```sql+postgres
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

```sql+sqlite
select
  id,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  created_at > datetime('now', '-30 days');
```