package goblog_test

import (
	"context"
	"goblog"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, tearDown := getTestDB(t)
	defer t.Cleanup(tearDown)

	ctx := context.Background()

	wantName := "Evanson Mwangi"
	wantEmail := "parish@github.io"

	u1, err := goblog.CreateUser(context.Background(), db, wantName, wantEmail)
	if err != nil {
		t.Fatalf("should be able to create a user")
	}

	u2, err := goblog.GetUserByEmail(ctx, db, wantEmail)

	if u2.Name != u1.Name {
		t.Fatalf("want name %s, got %s", u1.Name, u2.Name)
	}

	u3, err := goblog.GetUserByName(ctx, db, u1.Name)
	if err != nil {
		t.Fatal(err)
	}

	if u3.Email != u1.Email {
		t.Fatalf("want email %s, got %s", u1.Email, u3.Email)
	}
}
