---
organization: Turbot
category: ["asset management"]
icon_url: "/images/plugins/turbot/hubspot.svg"
brand_color: "#FF5C35"
display_name: "HubSpot"
short_name: "hubspot"
description: "Steampipe plugin to query contacts, deals, tickets and more from HubSpot."
og_description: "Query HubSpot with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/hubspot-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# HubSpot + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[HubSpot](https://www.hubspot.com/) is a CRM platform with all the software, integrations, and resources you need to connect marketing, sales, content management, and customer service. Each product in the platform is powerful on its own, but the real magic happens when you use them together.



List your HubSpot deals:

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

```
+-------------+---------------------------+----------+--------+----------+----------+----------------------+
| id          | created_at                | archived | amount | dealname | pipeline | dealstage            |
+-------------+---------------------------+----------+--------+----------+----------+----------------------+
| 13432979812 | 2023-05-24T17:09:39+05:30 | false    | 10000  | test     | default  | appointmentscheduled |
+-------------+---------------------------+----------+--------+----------+----------+----------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/hubspot/tables)**

## Quick start

### Install

Download and install the latest HubSpot plugin:

```sh
steampipe plugin install hubspot
```

### Credentials

| Item        | Description                                                                                                                                                                             |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | HubSpot requires a [Private App Token](https://developers.hubspot.com/docs/api/private-apps) for all requests.                                                                          |
| Permissions | The permission scope of Private App Tokens is set by the Admin at the creation time of the tokens.                                                                                      |
| Radius      | Each connection represents a single HubSpot Installation.                                                                                                                               |
| Resolution  | 1. Credentials explicitly set in a Steampipe config file (`~/.steampipe/config/hubspot.spc`)<br />2. Credentials specified in environment variables, e.g., `HUBSPOT_PRIVATE_APP_TOKEN`. |

Each table needs the private app token to carry the scope listed below. If a scope is missing, the table returns a permission error rather than an empty result.

| Tables                                                                            | Required scope                                                               |
| --------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `hubspot_contact`                                                                 | `crm.objects.contacts.read` (`crm.schemas.contacts.read` for custom properties)   |
| `hubspot_company`                                                                 | `crm.objects.companies.read` (`crm.schemas.companies.read` for custom properties) |
| `hubspot_deal`                                                                    | `crm.objects.deals.read` (`crm.schemas.deals.read` for custom properties)         |
| `hubspot_ticket`                                                                  | `tickets`                                                                    |
| `hubspot_owner`                                                                   | `crm.objects.owners.read`                                                    |
| `hubspot_blog_post`, `hubspot_cms_audit_log`                                      | `content`                                                                    |
| `hubspot_domain`                                                                  | `cms.domains.read`                                                           |
| `hubspot_hub_db`                                                                  | `hubdb`                                                                      |
| `hubspot_user`, `hubspot_user_role`                                               | `settings.users.read`                                                        |
| `hubspot_team`                                                                    | `settings.users.teams.read`                                                  |
| `hubspot_login_activity`, `hubspot_security_activity`                             | `account-info.security.read`                                                 |
| `hubspot_access_token`                                                            | None, the table introspects the token itself                                 |

### Configuration

Installing the latest hubspot plugin will create a config file (`~/.steampipe/config/hubspot.spc`) with a single connection named `hubspot`:

Configure your account details in `~/.steampipe/config/hubspot.spc`:

```hcl
connection "hubspot" {
  plugin = "hubspot"

  # The HubSpot Private APP Token. Required.
  # Get your Private APP token from HubSpot https://developers.hubspot.com/docs/api/private-apps.
  # Can also be set with the `HUBSPOT_PRIVATE_APP_TOKEN` environment variable.
  # private_app_token = "pat-na1-70271006-11d8-4a5d-9169-b12f4327e5b"
}
```

Alternatively, you can also use the standard HubSpot environment variable to obtain credentials **only if `private_app_token` is not specified** in the connection:

```sh
export HUBSPOT_PRIVATE_APP_TOKEN=pat-na1-70271006-11d8-4a5d-9169-b12f4327e5b
```


