## v0.1.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#11](https://github.com/turbot/steampipe-plugin-hubspot/pull/11))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#11](https://github.com/turbot/steampipe-plugin-hubspot/pull/11))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hubspot/blob/main/docs/LICENSE). ([#11](https://github.com/turbot/steampipe-plugin-hubspot/pull/11))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#10](https://github.com/turbot/steampipe-plugin-hubspot/pull/10))

## v0.0.2 [2023-11-17]

_Bug fixes_

- Fixed the plugin brand colour. 

## v0.0.1 [2023-11-17]

_What's new?_

- New tables added
  - [hubspot_blog_post](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_blog_post)
  - [hubspot_company](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_company)
  - [hubspot_contact](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_contact)
  - [hubspot_deal](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_deal)
  - [hubspot_domain](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_domain)
  - [hubspot_hub_db](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_hub_db)
  - [hubspot_owner](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_owner)
  - [hubspot_ticket](https://hub.steampipe.io/plugins/turbot/hubspot/tables/hubspot_ticket)
