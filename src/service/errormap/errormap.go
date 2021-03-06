package errormap

const (
	Success           = 0
	Exist             = 401
	NotExist          = 402
	PwdError          = 403
	ParamsError       = 405
	IllegalTS         = 406
	UnknowServerError = 500
)

var errorMap = map[int]string{
	401: "user exist",
	402: "user not exist",
	403: "pwd not correct",
	405: "params illegal",
	406: "illegal timestamp",
	500: "unknown server error",
}

func ErrorMsg(errCode int) string {
	return errorMap[errCode]
}
