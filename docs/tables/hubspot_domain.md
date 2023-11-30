---
title: "Steampipe Table: hubspot_domain - Query HubSpot Domains using SQL"
description: "Allows users to query Domains in HubSpot, providing specific details about each domain including its ID, name, created and updated timestamps, and more."
---

# Table: hubspot_domain - Query HubSpot Domains using SQL

HubSpot Domains are resources within the HubSpot platform that represent the domains associated with a particular HubSpot portal. These domains are used to host the content created within HubSpot, including landing pages, blog posts, and other forms of web content. Each domain is unique to a particular portal and contains specific details such as its ID, name, and the timestamps for when it was created and last updated.

## Table Usage Guide

The `hubspot_domain` table provides insights into the domains within HubSpot. As a web content manager or digital marketer, you can explore domain-specific details through this table, including the domain's ID, name, created and updated timestamps, and more. Use this table to maintain an inventory of your domains, track their creation and update times, and manage your web content more effectively.

## Examples

### Basic info
Explore the basic information of your Hubspot domains to understand their creation details and DNS configurations. This can help assess the correctness of the DNS and identify any potential issues with the domain settings.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain;
```

### List legacy domains
Discover the segments that still use legacy domains, enabling you to assess the elements within your system that may require updates or changes. This is particularly useful in maintaining system efficiency and staying up-to-date with newer domain standards.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  is_legacy;
```

### List domains where SSL is enabled
Explore domains that have SSL security enabled to ensure secure data transmission and improve website credibility. This is particularly beneficial for evaluating security measures and identifying potential improvements.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  is_ssl_enabled;
```

### List domains which are used for email
Uncover the details of domains that are configured for email usage. This is helpful for conducting audits or troubleshooting email delivery issues.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  is_used_for_email;
```

### List primary blog domains
Gain insights into the primary blog domains, focusing on their creation date, title, and email associated with them. This is beneficial for understanding the configuration and status of your blog domains, especially their DNS correctness and actual CNAME and IP.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  primary_blog;
```

### List domains that are not associated with any team
Discover the segments that consist of domains not linked to any team. This could be useful in identifying potential areas for team assignment or highlighting domains that may be under-utilized or overlooked.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  team_ids is null;
```

### List deletable domains
Explore which domains are marked as deletable in your Hubspot account. This can help maintain your domain list by identifying those that can be safely removed without disrupting your operations.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  deletable;
```

### List domains where setup is incomplete
Discover the segments that consist of domains with incomplete setup, enabling you to identify and address these areas to ensure all domains are fully operational. This is beneficial for maintaining a seamless and efficient digital infrastructure.

```sql
select
  id,
  title,
  created,
  primary_email,
  full_category_key,
  is_dns_correct,
  actual_cname,
  actual_ip
from
  hubspot_domain
where
  not is_setup_complete;
```