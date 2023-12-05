---
title: "Steampipe Table: hubspot_deal - Query HubSpot Deals using SQL"
description: "Allows users to query HubSpot Deals, providing insights into business transactions, deal stages, and related details."
---

# Table: hubspot_deal - Query HubSpot Deals using SQL

HubSpot Deals are a part of HubSpot's sales software, which allows businesses to track, organize, and manage deals or sales transactions. Deals are typically associated with stages and pipelines, and they help in visualizing and forecasting revenue, and in identifying bottlenecks in the sales process. HubSpot Deals provide a comprehensive view of the sales funnel, enabling businesses to make data-driven decisions.

## Table Usage Guide

The `hubspot_deal` table provides insights into business transactions managed through HubSpot's sales software. As a sales analyst or business manager, you can explore deal-specific details through this table, including deal stages, associated contacts, and forecasted revenue. Utilize it to uncover information about deals, such as their current status, associated pipeline, and potential bottlenecks in the sales process.

## Examples

### Basic info
Gain insights into the creation, archiving status, and financial details of various business deals. This can assist in tracking deal progress and managing sales pipelines effectively.

```sql+postgres
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

```sql+sqlite
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

### List deals that are not in the default pipeline
Discover the segments that are not part of the default pipeline in order to identify potential outliers or unique deals. This can be useful for revenue forecasting or sales strategy planning.

```sql+postgres
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

```sql+sqlite
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

### List closed deals
Discover the segments that have completed deals, allowing you to analyze past transactions and understand your sales performance. This is useful for assessing the effectiveness of your sales pipeline and identifying areas for improvement.

```sql+postgres
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

```sql+sqlite
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
  closedate <= date('now');
```

### List all archived deals
Uncover the details of all archived deals within your system to track their creation date, amount, name, pipeline, and stage. This allows for effective management and review of past business transactions.

```sql+postgres
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

```sql+sqlite
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
  archived = 1;
```

### List deals created in the last 30 days
Explore which deals were initiated in the past month. This can help businesses assess recent activity and understand the current state of their sales pipeline.

```sql+postgres
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

```sql+sqlite
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
  created_at > datetime('now', '-30 days');
```

### List new business deals
Explore which new business deals have been initiated by analyzing factors such as deal amount, stage, and whether it's archived. This can help in understanding the growth and progress of your business.

```sql+postgres
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

```sql+sqlite
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
Explore which deals have an amount exceeding $10,000 to gain insights into high-value transactions and assess their associated details such as deal name, pipeline, and deal stage. This could be beneficial in identifying lucrative opportunities and understanding the distribution of high-value deals across different stages.

```sql+postgres
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

```sql+sqlite
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
Explore which deals are in the appointment-scheduled stage, allowing you to assess the elements within your sales pipeline that are moving towards a potential close. This can help you manage your resources and prioritize your sales efforts effectively.

```sql+postgres
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

```sql+sqlite
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