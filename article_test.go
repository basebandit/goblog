package goblog_test

import (
	"context"
	"goblog"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateArticle(t *testing.T) {
	db, tearDown := getTestDB(t)
	t.Cleanup(tearDown)

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

}
