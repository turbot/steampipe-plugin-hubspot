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

### List deals which are new business

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
  dealtype = 'newbusiness';
```

### List deals where amount is more than $10000

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
  amount::int > 10000;
```

### List deals which are in appointmentscheduled stage

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
  dealstage = 'appointmentscheduled';
```