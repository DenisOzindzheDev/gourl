package response

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
