# Table: hubspot_deal

In HubSpot, a Deal represents a potential sales opportunity or a specific business transaction that your company is pursuing with a contact or a company. The Deal object is a core component of HubSpot's CRM (Customer Relationship Management) system and is used to track and manage the progress of sales deals from initial contact to closure.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal;
```

### List deals which are not in the default pipeline

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  pipeline <> 'default';
```

### List unclosed deals

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  closedate <= now();
```

### List all archived deals

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  archived;
```

### List deals created in the last 30 days

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  created_at > now() - interval '30 days';
```

### List deals which are a new business

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  dealtype = 'newbusiness';
```

### List deals where the amount is more than $10000

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  amount > 10000;
```

### List deals which are in an appointment-scheduled stage

```sql
select
  id,
  created_at,
  archived,
  amount,
  dealname,
  pipeline,
  dealstage
from
  hubspot_deal
where
  dealstage = 'appointmentscheduled';
```
