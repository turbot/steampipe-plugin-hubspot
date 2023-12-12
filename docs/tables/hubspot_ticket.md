---
title: "Steampipe Table: hubspot_ticket - Query Hubspot Tickets using SQL"
description: "Allows users to query Hubspot Tickets, specifically providing insights into ticket details such as status, priority, creation time, and associated contacts."
---

# Table: hubspot_ticket - Query Hubspot Tickets using SQL

Hubspot Tickets is a feature within the Hubspot Service Hub that allows users to track, prioritize, and solve customer support inquiries. It provides a centralized way to manage and respond to customer issues, ensuring efficient customer service. Hubspot Tickets helps businesses stay informed about customer issues and take appropriate actions when needed.

## Table Usage Guide

The `hubspot_ticket` table provides insights into tickets within Hubspot Service Hub. As a Customer Support Analyst, explore specific ticket details through this table, including status, priority, creation time, and associated contacts. Utilize it to track and manage customer issues, prioritize tasks, and ensure efficient customer service.

## Examples

### Basic info
Explore which tickets have been archived and their corresponding pipeline stages and priorities. This could be particularly useful for understanding the distribution of workload and identifying potential bottlenecks in your customer service process.

```sql+postgres
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

```sql+sqlite
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
Explore which tickets have been marked as high-priority in your system to focus your team's attention on the most crucial issues. This helps in efficient resource allocation and ensures timely resolution of critical matters.

```sql+postgres
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

```sql+sqlite
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
  hs_ticket_priority = 'HIGH';
```

### List all archived tickets
Explore which tickets have been archived to maintain an organized record and prioritize tasks based on the pipeline stage and ticket priority. This can help in tracking the progress of issues and understanding the efficiency of the ticket resolution process.

```sql+postgres
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

```sql+sqlite
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
Explore which tickets have been created in the last 30 days to assess recent customer issues and their priorities. This helps in identifying trends in customer concerns and managing resources effectively.

```sql+postgres
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

```sql+sqlite
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
  created_at > datetime('now', '-30 days');
```

### List tickets which are submitted via phone
Discover the segments that encompass tickets submitted through phone calls. This could be useful in understanding the nature and volume of phone-based customer interactions for strategic planning.

```sql+postgres
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

```sql+sqlite
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
Explore which customer support tickets are still open. This can help prioritize tasks and manage workload effectively in a customer service setting.

```sql+postgres
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

```sql+sqlite
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