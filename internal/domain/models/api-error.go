package models

type ApiError struct {
	Err  error
	Code int
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}
