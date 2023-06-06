# Table: hubspot_owner

In the context of HubSpot, the term Owner refers to the user responsible for managing and working on a specific object, such as a contact, company, deal, or task. The Owner property is a common attribute in HubSpot's CRM (Customer Relationship Management) system and is used to assign ownership and track responsibility for various records within your organization.

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
  last_name,
  user_id
from
  hubspot_owner;
```

### List all archived users

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

### List users created in the last 30 days

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

### List users who are not associated with any team

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
