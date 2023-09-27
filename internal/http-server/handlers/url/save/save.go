package save //save handler

import (
	"net/http"
	"url-shorner/internal/config"
	"url-shorner/lib/api/response"
	"url-shorner/lib/logger/sl"
	"url-shorner/lib/random"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	//_"github.com/go-playground/validator/v10"
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
		const op = "handlers.save.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
			//slog.String("url", r.URL.String()),
			//slog.String("method", r.Method),
		)
		var req Request
		err := render.DecodeJSON(r.Body, &req) //parse request body
		if err != nil {
			log.Error("error decoding request")
			render.JSON(w, r, response.Error("JSON decoding error")) // render error response
			return
		}
		log.Info("request body decoded", slog.Any("request", req)) //log request body with message

		if err := validator.New().Struct(req); err != nil { //validate request
			validateErr := err.(validator.ValidationErrors) // new validation errors obj
			log.Error("error validating request", sl.Err(err))
			render.JSON(w, r, response.Error("Invalid request"))      // render error response
			render.JSON(w, r, response.ValidationErrors(validateErr)) //render errors from response handler
			return                                                    // :D
		}
		var defaultAlliasLength int = int(config.MustLoad().DefaultAlliasLength)
		alias := req.Alias // alias
		if alias == "" {
			req.Alias = random.NewRandomString(defaultAlliasLength) // generate random alias
		}
	}
}
