package models

// Структуры таблиц баз данных.
type ResponseId struct {
	ID int64 `json:"id"`
}

type ResponseErr struct {
	Error string `json:"error"`
}

type Advert struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Contacts    string `json:"contacts"`
	Price       string `json:"price"`
	Removed     bool   `json:"removed"`
	CreatedAt   string `json:"createdAt"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phon  string `json:"phon"`
}

const (
	Schema_advert = `CREATE TABLE IF NOT EXISTS adverts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		author VARCHAR(128) NOT NULL DEFAULT '',
		contacts VARCHAR(128) NOT NULL DEFAULT '',
		price VARCHAR(128) NOT NULL DEFAULT '',
		removed BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX IF NOT EXISTS idxAdvertsId ON adverts (id);`

	Schema_user = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(128) NOT NULL DEFAULT '',
		email VARCHAR(128) NOT NULL DEFAULT '',
		phon VARCHAR(128) NOT NULL DEFAULT '');
		CREATE INDEX IF NOT EXISTS idxUsersId ON users (id);`
)
