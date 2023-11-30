---
title: "Steampipe Table: hubspot_blog_post - Query HubSpot Blog Posts using SQL"
description: "Allows users to query Blog Posts in HubSpot, providing insights into the content, performance, and metadata of each post."
---

# Table: hubspot_blog_post - Query HubSpot Blog Posts using SQL

HubSpot's Blog Post feature is a part of its broader content management system, allowing users to create, manage, and analyze blog content. It offers a platform for publishing rich, SEO-optimized content to engage audiences and drive traffic. Blog Posts in HubSpot are an essential tool for inbound marketing strategies, offering features like scheduling, analytics, and integrated CTAs.

## Table Usage Guide

The `hubspot_blog_post` table provides insights into the blog posts within HubSpot's content management system. As a content manager or marketing analyst, explore post-specific details through this table, including content, performance, and associated metadata. Utilize it to uncover information about posts, such as their SEO performance, engagement metrics, and the effectiveness of integrated CTAs.

## Examples

### Basic info
Explore the status and details of blog posts in your Hubspot account. This query can help you understand the distribution of your content, identify who's producing it, and when it's being published.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  archived,
  author_name,
  publish_date
from
  hubspot_blog_post;
```

### List all published blog posts
Explore all the blog posts that are currently live to understand the range of topics and authors contributing to your content. This can help in assessing the diversity of your content and identifying any gaps that need to be filled.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  archived,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  currently_published;
```

### List all archived blog posts
Discover the segments that contain all your archived blog posts. This aids in understanding the range and history of content that is no longer actively promoted but may still hold value.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  archived,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  archived;
```

### List blog posts created by a specific author
Explore which blog posts have been crafted by a specific author to gain insights into their productivity and content focus. This can be beneficial for content management and planning, allowing you to identify popular topics and authors.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  archived,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  author_name = 'John Doe';
```

### Get blog posts created by a specific owner
Explore blog posts authored by a specific individual to understand their contribution and publishing pattern. This can be particularly useful in content management scenarios, where tracking the work of specific authors is crucial for editorial oversight and planning.

```sql
select
  p.id,
  p.title,
  slug,
  campaign,
  state,
  author_name,
  publish_date
from
  hubspot_blog_post as p,
  hubspot_owner as o
where
  created_by_id::int = user_id
  and o.first_name = 'john';
```

### List blog posts in a specific category
Discover the segments that fall under a particular category within your blog posts. This can help you understand the distribution of your content and strategize your future posts accordingly.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  category_id = 123;
```

### List blog posts that have a featured image
Explore the blog posts that have a featured image to understand the impact of visual content on reader engagement. This can be useful in enhancing your content strategy by focusing on posts that are more likely to attract and retain readers.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  use_featured_image;
```

### List blog posts published in a specific campaign
Discover the segments that include blog posts published under a specific campaign. This can be useful for marketers aiming to analyze the performance of individual campaigns and their related content.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  campaign = 'CAMPAIGN123';
```

### Get blog posts that have public access rules enabled
Discover the segments that have public access rules enabled, allowing you to pinpoint blog posts that are accessible to the general public. This can be useful in assessing the visibility and reach of your content.

```sql
select
  id,
  title,
  slug,
  campaign,
  state,
  author_name,
  publish_date
from
  hubspot_blog_post
where
  public_access_rules_enabled;
```