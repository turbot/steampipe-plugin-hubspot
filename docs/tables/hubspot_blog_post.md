# Table: hubspot_blog_post

A HubSpot blog post is an article or piece of content that you create and publish on the HubSpot blogging platform. It allows you to share information, insights, and updates with your audience. HubSpot's blogging tools provide a user-friendly interface for writing, editing, and formatting your blog posts.

When creating a HubSpot blog post, you can:

- Write and format your content: Use the text editor to write your blog post, format text, add headings, apply styles, and insert images or media.

- Optimize for search engines: HubSpot provides SEO (search engine optimization) tools to optimize your blog post for better visibility in search engine results. You can add meta titles, descriptions, and keywords, and HubSpot offers recommendations to improve your SEO.

- Categorize and tag your blog post: Organize your blog posts into categories and assign relevant tags to make it easier for readers to find related content on your blog.

- Schedule or publish your blog post: Choose a publishing date and time to schedule your blog post in advance or publish it immediately.

- Analyze performance: HubSpot's analytics tools provide insights into your blog post's performance, including views, engagement, and conversion metrics.

HubSpot's blogging platform offers additional features like social sharing, commenting, and integrations with other marketing tools. It enables you to create and manage your blog content within the HubSpot ecosystem, integrating it with your broader marketing strategy.

## Examples

### Basic info

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
