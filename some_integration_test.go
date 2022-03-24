// +build integration

package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestArbitraryIntegration(t *testing.T) {
	// connect to postgres server (hardcoding connection string to work with CICD only for now)
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("Unable to open postgres connection: %s", err)
	}
	// wait for the test database to be available
	for attempt := 0; attempt < 5; attempt++ {
		if err := db.Ping(); err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	// time to use the (empty) database for testing
	assert.NoError(t, db.Ping())
}
