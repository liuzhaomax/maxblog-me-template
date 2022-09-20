package core

var ErrorMsg = map[int]string{
	996: "Upstream error",
	997: "Downstream error",
	998: "System internal error",
	999: "Unknown error",
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Err     error  `json:"error"`
}

func (err *Error) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func HandleError(errorCode int, err error) *Error {
	var errObj = new(Error)
	errObj.Code = errorCode
	errObj.Message = ErrorMsg[errorCode]
	if err != nil {
		errObj.Err = err
	}
	return errObj
}
