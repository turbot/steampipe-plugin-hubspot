# Table: hubspot_contact

In HubSpot, a Contact refers to an individual person or a potential customer with whom you have a relationship or interact within the context of your business. The Contact object is a fundamental component of HubSpot's CRM (Customer Relationship Management) system and serves as a central repository for storing and managing information related to your contacts, their interactions, activities, and more.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact;
```

### List all archived contacts

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  archived;
```

### List contacts created in the last 30 days

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  created_at > now() - interval '30 days';
```

### List contacts who are from Queensland

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  state = 'QLD';
```

### List lead contacts

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  lifecyclestage = 'lead';
```

### List contacts who have never closed any deal

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  recent_deal_close_date is null;
```

### List contacts who are Salespersons

```sql
select
  id,
  created_at,
  archived,
  email,
  firstname,
  lastname
from
  hubspot_contact
where
  jobtitle = 'Salesperson';
```
