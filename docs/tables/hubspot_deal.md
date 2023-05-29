# Table: hubspot_deal

In HubSpot, a Deal represents a potential sales opportunity or a specific business transaction that your company is pursuing with a contact or a company. The Deal object is a core component of HubSpot's CRM (Customer Relationship Management) system and is used to track and manage the progress of sales deals from initial contact to closure.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  archived,
  amount,
  deal_name,
  pipeline,
  deal_stage
from
  hubspot_deal;
```

### List deals which are not in default pipeline

```sql
select
  id,
  title,
  created_at,
  archived,
  amount,
  deal_name,
  pipeline,
  deal_stage
from
  hubspot_deal
where
  pipeline <> 'default';
```

### List unclosed deals

```sql
select
  id,
  title,
  created_at,
  archived,
  amount,
  deal_name,
  pipeline,
  deal_stage
from
  hubspot_deal
where
  close_date <= now();
```

### List all archived deals

```sql
select
  id,
  title,
  created_at,
  archived,
  amount,
  deal_name,
  pipeline,
  deal_stage
from
  hubspot_deal
where
  archived;
```

### List deals created in last 30 days

```sql
select
  id,
  title,
  created_at,
  archived,
  amount,
  deal_name,
  pipeline,
  deal_stage
from
  hubspot_deal
where
  created_at > now() - interval '30 days';
```

### Get deals associated with a specific company

```sql
select
  deal.id,
  deal.title,
  deal.created_at,
  deal.archived,
  deal.amount,
  deal.deal_name,
  deal.pipeline,
  deal.deal_stage
from
  hubspot_company as com,
  hubspot_deal as deal,
  jsonb_array_elements(associations_with_companies) as c
where
  c ->> 'id' = com.id
  and name = 'newCompany';
```

### Get deals associated with a specific contact

```sql
select
  deal.id,
  deal.title,
  deal.created_at,
  deal.archived,
  deal.amount,
  deal.deal_name,
  deal.pipeline,
  deal.deal_stage
from
  hubspot_contact as con,
  hubspot_deal as deal,
  jsonb_array_elements(associations_with_contacts) as c
where
  c ->> 'id' = com.id
  and first_name = 'Brian';
```

### Get deals associated with a specific ticket

```sql
select
  deal.id,
  deal.title,
  deal.created_at,
  deal.archived,
  deal.amount,
  deal.deal_name,
  deal.pipeline,
  deal.deal_stage
from
  hubspot_ticket as tik,
  hubspot_deal as deal,
  jsonb_array_elements(associations_with_tickets) as t
where
  t ->> 'id' = tik.id
  and subject = 'ticket_subjects';
```
