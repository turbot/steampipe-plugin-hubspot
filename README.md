![image](https://hub.steampipe.io/images/plugins/turbot/hubspot-social-graphic.png)

# HubSpot Plugin for Steampipe

Use SQL to query contacts, deals, tickets and more from HubSpot.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/hubspot)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/hubspot/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-hubspot/issues)

## Quick start

### Install

Download and install the latest HubSpot plugin:

```bash
steampipe plugin install hubspot
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/hubspot#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/hubspot#configuration).

Configure your account details in `~/.steampipe/config/hubspot.spc`:

```hcl
connection "hubspot" {
  plugin = "hubspot"

  # Authentication information
  private_app_token = "pat-na1-70271006-11d8-4a5d-9169-b12f4327e5b"
}
```

Or through environment variables:

```sh
export HUBSPOT_PRIVATE_APP_TOKEN=pat-na1-70271006-11d8-4a5d-9169-b12f4327e5b
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-hubspot.git
cd steampipe-plugin-hubspot
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/hubspot.spc
```

Try it!

```
steampipe query
> .inspect hubspot
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-hubspot/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hubspot/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [HubSpot Plugin](https://github.com/turbot/steampipe-plugin-hubspot/labels/help%20wanted)
