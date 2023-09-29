package storage

import "errors"

// errors
var (
	ErrURLNotFound    = errors.New("URL not found")
	ErrURLExists      = errors.New("URL already exists")
	ErrURLNoRows      = errors.New("URL has no rows")
	ErrDuplicateAlias = errors.New("Alias already exists")
	ErrAliasNotFound  = errors.New("Alias not found")
	ErrDuplicateURL   = errors.New("Duplicate URL")
)
