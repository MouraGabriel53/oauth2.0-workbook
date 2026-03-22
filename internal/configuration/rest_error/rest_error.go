package resterror

type RestError struct {
	Title  string  `json:"title"`
	Status int     `json:"status"`
	Detail string  `json:"detail,omitempty"`
	Causes []cause `json:"causes,omitempty"`
}

type cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (re *RestError) Error() string {
	return re.Detail
}

func NewInternalServerError(detail string) *RestError {
	return &RestError{
		Title:  "internal server error",
		Status: 500,
		Detail: detail,
	}
}
