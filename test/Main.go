package Test_Main

import (
	"errors"
	"gopurs/output/Promise.Internal"
	"gopurs/output/gopurs_runtime"
)

var HelloPromise = Promise_Internal.Resolve(gopurs_runtime.Box("Hello"))
var GoodbyePromise = Promise_Internal.Reject(gopurs_runtime.Box("Goodbye"))
var ErrPromise = Promise_Internal.Reject(gopurs_runtime.Box(errors.New("err")))
var CustomErrPromise = Promise_Internal.Reject(gopurs_runtime.Box(map[string]interface{}{"code": "err"}))
