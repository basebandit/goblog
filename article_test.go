package goblog_test

import (
	"context"
	"goblog"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateArticle(t *testing.T) {
	db, tearDown := getTestDB(t)
	defer t.Cleanup(tearDown)

	article := goblog.Article{
		Title:       "This is test article",
		Description: goblog.StringPointer("This is  test description"),
		Body:        " This is a test body",
	}

	ctx := context.Background()

	_, err := goblog.CreateUser(ctx, db, "parish", "parish@github.io")
	if err != nil {
		t.Fatalf("should be able to create a user: %v", err)
	}

	u, err := goblog.GetUserByEmail(ctx, db, "parish@github.io")
	if err != nil {
		t.Fatal(err)
	}
	a1, err := goblog.CreateArticle(ctx, db, &article, u.ID)
	if err != nil {
		t.Fatalf("should be able to create an article: %v", err)
	}

	ua := goblog.UpdateArticle{
		Title: goblog.StringPointer("This is another test article"),
		Body:  goblog.StringPointer("This is another test article body"),
	}

	if err = goblog.UpdateArticleByID(ctx, db, a1.ID, ua); err != nil {
		t.Fatal(err)
	}

	a2, err := goblog.GetArticleByTitle(ctx, db, *ua.Title)
	if err != nil {
		t.Fatal(err)
	}

	if a2.Title != *ua.Title {
		t.Fatalf("want title %s, got %s", a1.Title, *ua.Title)
	}

	if err = goblog.DeleteArticleByTitle(ctx, db, *ua.Title); err != nil {
		t.Fatal(err)
	}

	_, err = goblog.GetArticleByTitle(ctx, db, *ua.Title)

	if err != nil {
		if err.Error() != "article not found" {
			t.Fatalf("want error: article not found, got %v", err)
		}
	} else {
		t.Fatalf("want error, got %v", err)
	}
}
