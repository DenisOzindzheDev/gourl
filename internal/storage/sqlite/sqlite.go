package sqlite

import (
	"database/sql"
	"fmt"
	"url-shorner/internal/storage"

	"github.com/mattn/go-sqlite3"
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

// storage save url
func (s *Storage) SaveURL(alias string, urlToSave string) (int64, error) {
	const op = "storage.sqlite.SaveURL" // function name for error messages
	//todo add statement
	stmt, err := s.db.Prepare("INSERT INTO url(alias, url) VALUES(?,?)")
	if err != nil {
		return 0, fmt.Errorf("%s error: %s", op, err)
	}
	//execute the statement
	res, err := stmt.Exec(alias, urlToSave)
	//todo make more efficient
	if err != nil {
		if sqlLiteErr, ok := err.(sqlite3.Error); ok && sqlLiteErr.ExtendedCode == sqlite3.ErrConstraintUnique || sqlLiteErr.ExtendedCode == sqlite3.ErrNoExtended(sqlite3.ErrConstraint) {
			return 0, fmt.Errorf("%s error: %s", op, storage.ErrURLExists)
		}
	}
	//get the last inserted id
	id, err := res.LastInsertId() //last inserted id
	if err != nil {
		return 0, fmt.Errorf("%s error: %s", op, err)
	}
	return id, nil //return the last inserted id and nil error
}

// clear database
func (s *Storage) TruncateDB() error {
	const op = "storage.sqlite.TruncateDB" // function name for error messages
	//todo add statement
	stmt, err := s.db.Prepare("DELETE FROM url")
	if err != nil {
		return fmt.Errorf("%s error: %s", op, err)
	}
	//execute the statement
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s error: %s", op, err)
	}
	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL" // function name for error messages
	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias =?")
	if err != nil {
		return "", fmt.Errorf("%s error: %s", op, err)
	}
	//execute the statement
	row := stmt.QueryRow(alias)
	var url string
	err = row.Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("%s error: %s", op, storage.ErrURLNotFound)
		}
		return "", fmt.Errorf("%s error: %s", op, err)
	}
	return url, nil
}
