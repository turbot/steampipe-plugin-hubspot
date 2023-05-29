# Table: hubspot_ticket

In the context of customer support or service management, a Ticket refers to a record or a request created by a customer or user to report an issue, seek assistance, or make an inquiry. Tickets are commonly used in helpdesk systems and customer support platforms, including HubSpot Service Hub, to efficiently manage and track customer interactions.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  archived,
  content,
  pipeline,
  pipeline_stage,
  ticket_priority
from
  hubspot_ticket;
```

### List high priority tickets

```sql
select
  id,
  title,
  created_at,
  archived,
  content,
  pipeline,
  pipeline_stage,
  ticket_priority
from
  hubspot_ticket
where
  ticket_priority = 'HIGH';
```

### List all archived tickets

```sql
select
  id,
  title,
  created_at,
  archived,
  content,
  pipeline,
  pipeline_stage,
  ticket_priority
from
  hubspot_ticket
where
  archived;
```

### List tickets created in last 30 days

```sql
select
  id,
  title,
  created_at,
  archived,
  content,
  pipeline,
  pipeline_stage,
  ticket_priority
from
  hubspot_ticket
where
  created_at > now() - interval '30 days';
```

### Get tickets associated with a specific company

```sql
select
  tik.id,
  tik.title,
  tik.created_at,
  tik.archived,
  tik.content,
  tik.pipeline,
  tik.pipeline_stage,
  tik.ticket_priority
from
  hubspot_company as com,
  hubspot_ticket as tik,
  jsonb_array_elements(associations_with_companies) as c
where
  c ->> 'id' = com.id
  and name = 'newCompany';
```

### Get tickets associated with a specific contact

```sql
select
  tik.id,
  tik.title,
  tik.created_at,
  tik.archived,
  tik.content,
  tik.pipeline,
  tik.pipeline_stage,
  tik.ticket_priority
from
  hubspot_contact as con,
  hubspot_ticket as tik,
  jsonb_array_elements(associations_with_contacts) as c
where
  c ->> 'id' = con.id
  and first_name = 'Brian';
```

### Get tickets associated with a specific deal

```sql
select
  tik.id,
  tik.title,
  tik.created_at,
  tik.archived,
  tik.content,
  tik.pipeline,
  tik.pipeline_stage,
  tik.ticket_priority
from
  hubspot_deal as deal,
  hubspot_ticket as tik,
  jsonb_array_elements(associations_with_deals) as d
where
  d ->> 'id' = deal.id
  and deal_name = 'final_deal';
```
