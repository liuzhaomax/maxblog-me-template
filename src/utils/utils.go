package utils

import (
	"runtime"
	"strconv"
)

func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	return function.Name()
}

func Str2Uint32(str string) (uint32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}
