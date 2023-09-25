package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// db structure
type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New" // function name for error messages
	//open database file
	db, err := sql.Open("sqlite3", storagePath) //path to database file
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", op, err)
	}
	//db statement stmt is statement
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alias TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", op, err)
	}
	//execute the statement
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", op, err)
	}
	return &Storage{db: db}, nil
}
