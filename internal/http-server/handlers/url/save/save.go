package save //save handler

import (
	"net/http"
	"url-shorner/lib/api/response"

	"golang.org/x/exp/slog"
)

// request handler structure
type Request struct {
	URL   string `json:"url" validate:"required,url"` //request in json format
	Alias string `json:"alias" omitempty:"true"`      //requaries this parameters
}
type Response struct {
	Response response.Response `json:"response" validate:"required"`
	Alias    string            `json:"alias" omitempty:"true"` //response alias (new alias)
}

// signature for the URLSaver interface
type URLSaver interface {
	SaveURL(alias string, urlToSave string) (int64, error) // method to save a URL interface #in storage/sqlite/sqlite.go
	//func (s *Storage) SaveURL(alias string, urlToSave string) (int64, error) {
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
