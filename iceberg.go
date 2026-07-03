// Package dataiceberg is a togo data backend that queries Apache Iceberg tables
// through DuckDB (which reads Iceberg natively). Select with
// `togo provider:use data iceberg`.
//
//   ICEBERG_DUCKDB_BIN  duckdb binary (default "duckdb")
//   ICEBERG_INIT        setup SQL run before each query
//                       (default "INSTALL iceberg; LOAD iceberg;")
//
// Point queries at your tables with iceberg_scan('s3://…/table') per DuckDB's
// iceberg extension.
package dataiceberg

import (
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/togo-framework/data"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("data-iceberg", togo.PriorityService+1, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapData, "iceberg", newIceberg(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "data-iceberg")
		}
		return nil
	})
}

type iceberg struct {
	bin  string
	init string
}

func newIceberg(k *togo.Kernel) *iceberg {
	return &iceberg{
		bin:  providers.Value(k, providers.CapData, "iceberg", "duckdb_bin", "duckdb", false),
		init: providers.Value(k, providers.CapData, "iceberg", "init", "INSTALL iceberg; LOAD iceberg;", false),
	}
}

// assemble combines the init SQL with the query into one DuckDB script.
func (i *iceberg) assemble(query string) string {
	init := strings.TrimSpace(i.init)
	if init != "" && !strings.HasSuffix(init, ";") {
		init += ";"
	}
	return init + " " + query
}

func (i *iceberg) Query(ctx context.Context, query string, _ ...any) ([]data.Row, error) {
	out, err := exec.CommandContext(ctx, i.bin, "-json", "-c", i.assemble(query)).Output()
	if err != nil {
		return nil, err
	}
	var rows []data.Row
	if len(strings.TrimSpace(string(out))) == 0 {
		return nil, nil
	}
	if err := json.Unmarshal(out, &rows); err != nil {
		return nil, err
	}
	return rows, nil
}
