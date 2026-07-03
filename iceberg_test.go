package dataiceberg

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/togo-framework/togo"
)

func TestAssemble(t *testing.T) {
	i := newIceberg(togo.New())
	got := i.assemble("SELECT 1")
	if !strings.Contains(got, "LOAD iceberg;") || !strings.HasSuffix(got, "SELECT 1") {
		t.Fatalf("assemble = %q", got)
	}
}

func TestQueryParsesDuckDBJSON(t *testing.T) {
	dir := t.TempDir()
	fake := filepath.Join(dir, "duckdb")
	// fake duckdb: emits a JSON array of rows regardless of args
	if err := os.WriteFile(fake, []byte("#!/bin/sh\necho '[{\"n\":42,\"city\":\"riyadh\"}]'\n"), 0755); err != nil {
		t.Fatal(err)
	}
	i := &iceberg{bin: fake, init: "LOAD iceberg;"}
	rows, err := i.Query(context.Background(), "SELECT * FROM iceberg_scan('s3://x/t')")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0]["city"] != "riyadh" {
		t.Fatalf("rows = %v", rows)
	}
}
