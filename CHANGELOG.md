## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#37](https://github.com/turbot/steampipe-plugin-hubspot/pull/37))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#37](https://github.com/turbot/steampipe-plugin-hubspot/pull/37))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#29](https://github.com/turbot/steampipe-plugin-hubspot/pull/29))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#29](https://github.com/turbot/steampipe-plugin-hubspot/pull/29))

## v0.1.1 [2024-02-26]

_Bug fixes_

- Fixed the plugin to return nil instead of an error when API credentials are not set in the `*.spc` file. ([#14](https://github.com/turbot/steampipe-plugin-hubspot/pull/14))
- Updated the type of the `default` column for the dynamic columns from `JSON` to `string`. ([#16](https://github.com/turbot/steampipe-plugin-hubspot/pull/16))

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
