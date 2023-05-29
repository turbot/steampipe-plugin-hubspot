# Table: hubspot_company

In HubSpot, the Company object represents a company or organization with which you interact or have a business relationship. It is a core entity in HubSpot's CRM (Customer Relationship Management) system and serves as a central repository for storing and managing information related to your company contacts, deals, interactions, and more.

## Examples

### Basic info

```sql
select
  id,
  title,
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
  title,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  archived;
```

### List companies created in last 30 days

```sql
select
  id,
  title,
  created_at,
  archived,
  domain,
  updated_at
from
  hubspot_company
where
  created_at > now() - interval '30 days';
```

### Get companies associated with a specific contact

```sql
select
  com.id,
  com.title,
  com.created_at,
  com.archived,
  com.domain,
  com.updated_at
from
  hubspot_contact as con,
  hubspot_company as com,
  jsonb_array_elements(associations_with_contacts) as c
where
  c ->> 'id' = con.id
  and first_name = 'Brian';
```

### Get companies associated with a specific deal

```sql
select
  com.id,
  com.title,
  com.created_at,
  com.archived,
  com.domain,
  com.updated_at
from
  hubspot_deal as deal,
  hubspot_company as com,
  jsonb_array_elements(associations_with_deals) as d
where
  d ->> 'id' = deal.id
  and deal_name = 'final_deal';
```

### Get companies associated with a specific ticket

```sql
select
  com.id,
  com.title,
  com.created_at,
  com.archived,
  com.domain,
  com.updated_at
from
  hubspot_ticket as tik,
  hubspot_company as com,
  jsonb_array_elements(associations_with_tickets) as t
where
  t ->> 'id' = tik.id
  and subject = 'ticket_subjects';
```