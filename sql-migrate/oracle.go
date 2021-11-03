//go:build oracle
// +build oracle

package main

import (
	migrate "github.com/aman00323/sql-migrate"
	_ "github.com/mattn/go-oci8"
)

func init() {
	dialects["oci8"] = migrate.OracleDialect{}
}
