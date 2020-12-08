package goblog_test

import (
	"context"
	"goblog"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, tearDown := getTestDB(t)
	defer t.Cleanup(tearDown)

	wantName := "Evanson Mwangi"
	wantEmail := "parish@github.io"

	_, err := goblog.CreateUser(context.Background(), db, wantName, wantEmail)
	if err != nil {
		t.Fatalf("should be able to create a user")
	}

	u2, err := goblog.GetUserByEmail(context.Background(), db, name)

	if u2.Name == wantName {
		t.Fatalf("want name %s, got %s", wantName, u2.Name)
	}

}
