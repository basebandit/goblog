package main

import (
	"fmt"
	"goblog"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	cfg := goblog.Config{
		Name:       os.Getenv("PG_DB"),
		Host:       os.Getenv("PG_HOST"),
		Port:       os.Getenv("PG_PORT"),
		User:       os.Getenv("PG_USER"),
		Password:   os.Getenv("PG_PASSWORD"),
		DisableTLS: true,
	}

	fmt.Printf("%+v\n", cfg)

	db, err := goblog.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := goblog.Migrate(db); err != nil {
		log.Fatalf("migrating error: %s", err)
	}

	fmt.Printf("waiting for the database to be ready...\n")

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
		log.Fatalf("database never ready: %v", pingError)
	}

	fmt.Println("connected to database successfully!")
}
