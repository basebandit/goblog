package goblog_test

import (
	"goblog"
	"log"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	name       = "goblog"
	host       = "localhost"
	port       = "5432"
	user       = "gb_user"
	password   = "gb_s3cr37"
	disableTLS = true
)

func getTestDB(t *testing.T) (*sqlx.DB, func()) {
	cfg := goblog.Config{
		Name:       name,
		Host:       host,
		Port:       port,
		User:       user,
		Password:   password,
		DisableTLS: true,
	}

	db, err := goblog.Connect(cfg)
	if err != nil {
		t.Fatalf("opening database connection: %v", err)
	}

	//Wait for the database to be ready. Wait 100ms longer between each attempt.
	//Do not try more than 20 times..
	var pingError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}

	if pingError != nil {
		t.Fatalf("database never ready: %v", pingError)
	}

	if err := goblog.Migrate(db); err != nil {
		log.Fatalf("migrating error: %s", err)
	}

	tearDown := func() {
		err := goblog.Drop(db)
		if err != nil {
			t.Fatalf("failed to drop test database: %v", err)
		}
	}

	return db, tearDown
}
