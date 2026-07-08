package model

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *AppError) Error() string {
	return err.Message
}
