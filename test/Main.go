package Test_Main

import (
	"errors"
	"gopurs/output/gopurs_runtime"
)

var HelloPromise = Promise_Internal_Resolve(gopurs_runtime.Box("Hello"))
var GoodbyePromise = Promise_Internal_Reject(gopurs_runtime.Box("Goodbye"))
var ErrPromise = Promise_Internal_Reject(gopurs_runtime.Box(errors.New("err")))
var CustomErrPromise = Promise_Internal_Reject(gopurs_runtime.Box(map[string]interface{}{"code": "err"}))
