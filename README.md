<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">data-iceberg</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a></sub></p>

A togo **data** backend that queries **Apache Iceberg** tables through **DuckDB**
(which reads Iceberg natively). Registers into
[data](https://github.com/togo-framework/data)'s slot.

```bash
togo install togo-framework/data-iceberg
togo provider:use data iceberg
# then query with DuckDB's iceberg_scan('s3://…/table')
```

| Key | Default | Purpose |
|---|---|---|
| `ICEBERG_DUCKDB_BIN` | `duckdb` | duckdb binary |
| `ICEBERG_INIT` | `INSTALL iceberg; LOAD iceberg;` | setup SQL |

> This is the "tier-out" escape hatch from `data-pg`: hot data stays in Postgres,
> cold/shared data lives in Iceberg on object storage — both queryable.

MIT © fadymondy
