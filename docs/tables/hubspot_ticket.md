# Table: hubspot_ticket

In the context of customer support or service management, a Ticket refers to a record or a request created by a customer or user to report an issue, seek assistance, or make an inquiry. Tickets are commonly used in helpdesk systems and customer support platforms, including HubSpot Service Hub, to efficiently manage and track customer interactions.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket;
```

### List high-priority tickets

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket;
where
  hs_ticket_priority = 'HIGH';
```

### List all archived tickets

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket
where
  archived;
```

### List tickets created in the last 30 days

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket
where
  created_at > now() - interval '30 days';
```

### List tickets which are submitted via phone

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket
where
  source_type = 'PHONE';
```

### List open tickets

```sql
select
  id,
  created_at,
  archived,
  content,
  hs_pipeline,
  hs_pipeline_stage,
  hs_ticket_priority
from
  hubspot_ticket
where
  time_to_close is null;
```
