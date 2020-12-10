package goblog

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Article is a blog entry
type Article struct {
	//ID is a unique identifier for each article
	ID string
	//Slug is a safe-shortened-format of the url
	Slug string `db:"slug" json:"slug"`
	//Title is the article's title
	Title string `db:"title" json:"title"`
	//Description gives an overview of the article's content. It is optional.
	Description *string `db:"description" json:"desc"`
	//Body is the article's content
	Body string `db:"body" json:"body"`
	//Author is the user who wrote/published the article
	Author string `json:"author"`
	//CreatedAt time the article was created.
	CreatedAt time.Time `db:"created_at" json:"created"`
	//UpdatedAt time the article was modified.
	UpdatedAt time.Time `db:"updated_at" json:"updated"`
}

//UpdateArticle defines article fields that can change
type UpdateArticle struct {
	Title       *string   `db:"title"`
	Description *string   `db:"description"`
	Body        *string   `db:"body"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func slugify(title string) string {
	return strings.ReplaceAll(title, " ", "-")
}

//StringPointer returns a reference to a string variable.
func StringPointer(s string) *string {
	str := s
	return &str
}

//CreateArticle creates a new article entry in the database.
func CreateArticle(ctx context.Context, db *sqlx.DB, a *Article, userID string) (Article, error) {

	articleID, err := uuid.NewV4()

	if err != nil {
		return Article{}, fmt.Errorf("failed to generate article id: %v", err)
	}

	article := Article{
		ID:        articleID.String(),
		Author:    userID,
		Body:      a.Body,
		Slug:      slugify(a.Title),
		Title:     a.Title,
		CreatedAt: time.Now(),
	}

	if a.Description != nil {
		article.Description = a.Description
	}

	const q = `insert into articles (id,author,title,slug,body,description,created_at) values ($1,$2,$3,$4,$5,$6,$7)`

	if _, err := db.ExecContext(ctx, q, articleID, userID, article.Title, article.Slug, article.Body, article.Description, article.CreatedAt); err != nil {
		return Article{}, errors.Wrap(err, "inserting article")
	}

	return article, nil
}

//GetArticleByTitle retrieves an article by title.
func GetArticleByTitle(ctx context.Context, db *sqlx.DB, title string) (Article, error) {
	const q = "select id,title,slug,description,body,created_at from articles where title=$1"

	var a Article
	if err := db.GetContext(ctx, &a, q, title); err != nil {
		if err == sql.ErrNoRows {
			return Article{}, errors.New("article not found")
		}
		return Article{}, errors.Wrapf(err, "selecting article %q", title)
	}

	return a, nil
}

//UpdateArticleByID updates an existing article entry
func UpdateArticleByID(ctx context.Context, db *sqlx.DB, id string, ua UpdateArticle) error {
	var a Article

	if ua.Title != nil {
		a.Title = *ua.Title
	}

	if ua.Body != nil {
		a.Body = *ua.Body
	}

	if ua.Description != nil {
		a.Description = ua.Description
	}

	const q = `UPDATE articles SET
	"title" = $2,
	"body" = $3,
	"description" = $4,
	"updated_at" = $5
	WHERE id = $1`

	if _, err := db.ExecContext(ctx, q, id, a.Title, a.Body, a.Description, time.Now()); err != nil {
		return errors.Wrap(err, "updating user")
	}

	return nil
}

//DeleteArticleByTitle deletes an article. Note the query does not cascade to the owner of the foreign key.
func DeleteArticleByTitle(ctx context.Context, db *sqlx.DB, title string) error {
	const q = `delete from articles where title=$1`

	if _, err := db.ExecContext(ctx, q, title); err != nil {
		return errors.Wrap(err, "deleting article")
	}

	return nil
}
