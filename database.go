package goblog

import (
	"fmt"
	"net"
	"net/url"

	"github.com/jmoiron/sqlx"
)

//Config defines configuration requirements of the database.
type Config struct {
	User       string
	Password   string
	Host       string
	Name       string
	Port       string
	DisableTLS bool
}

//Connect opens a connection to the database. Returns an open connection handle.
func Connect(cfg Config) (*sqlx.DB, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     net.JoinHostPort(cfg.Host, cfg.Port),
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())

}

//Migrate creates the tables in the database.
func Migrate(db *sqlx.DB) error {
	for _, q := range migrate {
		_, err := db.Exec(q)
		if err != nil {
			return fmt.Errorf("failed to migrate: %v; query: %q", err, q)
		}
	}
	return nil
}

//Drop reverses the effects of migrate.
func Drop(db *sqlx.DB) error {
	for _, q := range drop {
		_, err := db.Exec(q)
		if err != nil {
			return fmt.Errorf("failed to drop: %v;query: %q", err, q)
		}
	}
	return nil
}
