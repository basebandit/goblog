package goblog

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//User is the owner of the blog
type User struct {
	//ID is a unique identifier for each owner
	ID string `db:"id" json:"id"`
	//Name is the unique name of the owner
	Name string `db:"name" json:"name"`
	//Email is the owner's unique email address
	Email string `db:"email" json:"email"`
	//CreatedAt time this user was created
	CreatedAt time.Time `db:"created_at" json:"created"`
	//UpdatedAt time this user was updated
	UpdatedAt time.Time `db:"udated_at" json:"updated"`
}

//CreateUser creates a new user entry in the database.
func CreateUser(ctx context.Context, db *sqlx.DB, name, email string) (User, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return User{}, fmt.Errorf("failed to generate user id: %v", err)
	}

	u := User{
		ID:        id.String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	const q = `insert into users (id,name,email,created_at) values ($1,$2,$3,$4)`

	if _, err := db.ExecContext(ctx, q, u.ID, u.Name, u.Email, u.CreatedAt); err != nil {
		return User{}, errors.Wrap(err, "inserting user")
	}

	return u, nil
}
