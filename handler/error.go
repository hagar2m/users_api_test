package handler

type HTTPError struct {
	Status  int    // HTTP status code
	Message string // Error message
}

func (e *HTTPError) Error() string {
	return e.Message
}
