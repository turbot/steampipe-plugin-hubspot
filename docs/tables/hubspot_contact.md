---
title: "Steampipe Table: hubspot_contact - Query HubSpot Contacts using SQL"
description: "Allows users to query HubSpot Contacts, providing insights into contact details and interactions."
---

# Table: hubspot_contact - Query HubSpot Contacts using SQL

HubSpot Contacts is a resource within the HubSpot CRM that provides a comprehensive view of all interactions with a particular contact. It includes information such as contact details, communication history, and any associated deals or tasks. This resource helps businesses to maintain a centralized, up-to-date record of all contact interactions, enabling personalized and targeted communication strategies.

## Table Usage Guide

The `hubspot_contact` table provides insights into contact details and interactions within HubSpot CRM. As a Sales or Marketing professional, explore contact-specific details through this table, including communication history, associated deals, and tasks. Utilize it to uncover information about contacts, such as their engagement with your business, their preferences, and the effectiveness of your communication strategies.

## Examples

### Basic info
Explore which contacts have been archived in your Hubspot database. This can help you identify instances where contact details may need to be updated or restored, providing a practical tool for maintaining your contact list.

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
Explore which contacts have been archived in your Hubspot account. This can be useful for cleaning up your contact list or identifying potential leads that have been overlooked.

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
Determine the contacts that have been added within the past month. This allows you to keep track of recent additions to your network and ensure you are staying up-to-date with new connections.

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

### List contacts from Queensland
Explore which contacts are based in Queensland to target marketing campaigns more effectively. This helps in personalizing communication and improving customer engagement.

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
Discover the segments that consist of lead contacts in your Hubspot account. This query is useful in identifying prospective customers for targeted marketing or sales follow-ups.

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

### List contacts that have never closed any deal
Determine the areas in which contacts have not yet successfully closed any deals. This can be useful for identifying potential opportunities for engagement or training needs.

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
Explore which contacts in your database are designated as Salespersons. This can be beneficial in identifying potential leads or understanding your interaction with this specific job role.

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