# Table: hubspot_domain

In HubSpot, the Domain refers to the website domain associated with a company or contact. It represents the web address or URL of the company's website or the individual contact's website, if applicable.

The Domain property is commonly used to associate companies or contacts with their corresponding websites, allowing you to track and manage information related to online presence, website activity, and interactions. By capturing the domain information, you can gain insights into a company or contact's online footprint, website traffic, and engagement.

In HubSpot's CRM (Customer Relationship Management) system, you can store and utilize the Domain property to segment and categorize companies or contacts based on their website domains. This enables targeted marketing, personalized communication, and a better understanding of your audience's online behavior.

## Examples

### Basic info

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

### List domains that are primary blog domain

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