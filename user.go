package goblog

import (
	"time"
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
