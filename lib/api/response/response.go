package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// common data model for response
type Response struct {
	Status string `json:"status"`                 //200 ok 500 error
	Error  string `json:"error" omitempty:"true"` //error message
}

const (
	StatusOK    = "ok"
	StatusERROR = "error"
	//todo update statuses
)

func OK() Response {
	return Response{Status: StatusOK}
}
func Error(err string) Response {
	return Response{Status: StatusERROR, Error: err}
}

// get validation error and handle response
func ValidationErrors(err validator.ValidationErrors) Response {
	var errMsgs []string // error messages slice
	// for range of all errors
	for _, e := range err {
		switch e.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required filed", e.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not valid url", e.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not valid filed", e.Field()))
		}
	}
	// get all errors in array
	return Response{
		Status: StatusERROR,
		Error:  strings.Join(errMsgs, " | "), // join all error messages
	}
}
