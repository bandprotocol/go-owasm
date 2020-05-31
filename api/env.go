package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "bindings.h"
import "C"
import "unsafe"

type Env struct {
	AskCount int
}

//export cGetAskCount
func cGetAskCount(e *C.env_t) C.int64_t {
	env := *(*Env)(unsafe.Pointer(e))
	return C.int64_t(env.AskCount)
}
