package errors

type LibmaError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Code   int    `json:"code"`
}

var LibmaErrors = map[string]LibmaError{
	"BadRequest": {
		Status: 400,
		Msg:    "Bad Request",
		Code:   400000,
	},
	"InternalServerError": {
		Status: 500,
		Msg:    "Internal Server Error",
		Code:   500000,
	},
}

func GetLibmaError(err string) LibmaError {
	return LibmaErrors[err]
}

func GetLibmaErrorByCode(code int) LibmaError {
	for _, v := range LibmaErrors {
		if v.Code == code {
			return v
		}
	}
	return LibmaError{}
}
