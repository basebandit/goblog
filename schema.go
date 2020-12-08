package goblog

var migrate = []string{
	`
	CREATE TABLE IF NOT EXISTS users (
		id         UUID  PRIMARY KEY,
		name       TEXT NOT NULL,
		email      TEXT  NOT NULL,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ 
	);

	create unique index on users(lower(name));
	create unique index on users(email);

	CREATE TABLE IF NOT EXISTS tags (
		id         UUID    PRIMARY KEY,
		name       TEXT         NOT NULL,
		created_at TIMESTAMPTZ  NOT NULL
	);

	CREATE TABLE IF NOT EXISTS articles (
		id UUID PRIMARY KEY,
		title TEXT NOT NULL,
		author UUID,
		slug TEXT,
		body TEXT NOT NULL,
		description TEXT,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		FOREIGN KEY (author) REFERENCES users(id) ON DELETE CASCADE
	);
	create unique index on articles(title);
	create unique index on articles(slug);`,
}

var drop = []string{
	`drop table if exists users cascade`,
	`drop table if exists tags cascade`,
	`drop table if exists articles cascade`,
}
