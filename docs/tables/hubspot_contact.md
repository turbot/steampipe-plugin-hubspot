# Table: hubspot_contact

In HubSpot, a Contact refers to an individual person or a potential customer with whom you have a relationship or interact within the context of your business. The Contact object is a fundamental component of HubSpot's CRM (Customer Relationship Management) system and serves as a central repository for storing and managing information related to your contacts, their interactions, activities, and more.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name
from
  hubspot_contact;
```

### List all archived contacts

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name
from
  hubspot_contact
where
  archived;
```

### List contacts created in last 30 days

```sql
select
  id,
  title,
  created_at,
  archived,
  email,
  first_name,
  last_name
from
  hubspot_contact
where
  created_at > now() - interval '30 days';
```

### Get contacts associated with a specific company

```sql
select
  con.id,
  con.title,
  con.created_at,
  con.archived,
  con.email,
  con.first_name,
  con.last_name
from
  hubspot_company as com,
  hubspot_contact as con,
  jsonb_array_elements(associations_with_companies) as c
where
  c ->> 'id' = com.id
  and name = 'newCompany';
```

### Get contacts associated with a specific deal

```sql
select
  con.id,
  con.title,
  con.created_at,
  con.archived,
  con.email,
  con.first_name,
  con.last_name
from
  hubspot_deal as deal,
  hubspot_contact as con,
  jsonb_array_elements(associations_with_deals) as d
where
  d ->> 'id' = deal.id
  and deal_name = 'final_deal';
```

### Get contacts associated with a specific ticket

```sql
select
  con.id,
  con.title,
  con.created_at,
  con.archived,
  con.email,
  con.first_name,
  con.last_name
from
  hubspot_ticket as tik,
  hubspot_contact as con,
  jsonb_array_elements(associations_with_tickets) as t
where
  t ->> 'id' = tik.id
  and subject = 'ticket_subjects';
```