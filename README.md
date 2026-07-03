<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/data-iceberg</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/data-iceberg"><img src="https://pkg.go.dev/badge/github.com/togo-framework/data-iceberg.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/data-iceberg
```

<!-- /togo-header -->

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
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
