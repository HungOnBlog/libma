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

	// Borrower
	// Detail code (code) of borrower errors is in range xxx100 - xxx199
	// Borrower errors -- 403100 - 403199
	"BorrowerLoginFailed": {
		Status: 403,
		Msg:    "Borrower login failed. Email and password do not match or borrower does not exist",
		Code:   403100,
	},
	// Borrower errors -- 404100 - 404199
	"BorrowerNotFound": {
		Status: 404,
		Msg:    "Borrower not found",
		Code:   404100,
	},
	// Borrower errors -- 409100 - 409199
	"BorrowerAlreadyExists": {
		Status: 409,
		Msg:    "Borrower (using this name or email) already exists. Please try again with another name or email",
		Code:   409100,
	},
	// Borrower errors -- 400100 - 400199
	"InvalidBorrowerCreationInfo": {
		Status: 400,
		Msg:    "Invalid borrower creation info. Please check 'name', 'email' and 'password' fields",
		Code:   400100,
	},
	"InvalidBorrowerUpdateInfo": {
		Status: 400,
		Msg:    "Invalid borrower update info. Please check 'name', 'email' and 'password' fields",
		Code:   400101,
	},
	"InvalidBorrowerLoginInfo": {
		Status: 400,
		Msg:    "Invalid borrower login info. Please check 'email' and 'password' fields",
		Code:   400102,
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
