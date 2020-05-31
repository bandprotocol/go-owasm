package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "bindings.h"
import "C"
import (
	"unsafe"
)

type EnvInterface interface {
	GetCalldata() []byte
	SetReturnData([]byte)
	GetAskCount() int64
	GetMinCount() int64
	GetAnsCount() int64
	AskExternalData(eid int64, did int64, data []byte)
	GetExternalData(eid int64, vid int64) []byte
}

type envIntl struct {
	ext      EnvInterface
	calldata C.Span
	extData  map[[2]int64]C.Span
}

func createEnvIntl(ext EnvInterface) *envIntl {
	return &envIntl{
		ext:      ext,
		calldata: copySpan(ext.GetCalldata()),
		extData:  make(map[[2]int64]C.Span),
	}
}

func destroyEnvIntl(e *envIntl) {
	freeSpan(e.calldata)
	for _, span := range e.extData {
		freeSpan(span)
	}
}

//export cGetCalldata
func cGetCalldata(e *C.env_t) C.Span {
	return (*(*envIntl)(unsafe.Pointer(e))).calldata
}

//export cSetReturnData
func cSetReturnData(e *C.env_t, span C.Span) {
	(*(*envIntl)(unsafe.Pointer(e))).ext.SetReturnData(readSpan(span))
}

//export cGetAskCount
func cGetAskCount(e *C.env_t) C.int64_t {
	return C.int64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetAskCount())
}

//export cGetMinCount
func cGetMinCount(e *C.env_t) C.int64_t {
	return C.int64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetMinCount())
}

//export cGetAnsCount
func cGetAnsCount(e *C.env_t) C.int64_t {
	return C.int64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetAnsCount())
}

//export cAskExternalData
func cAskExternalData(e *C.env_t, eid C.int64_t, did C.int64_t, span C.Span) {
	(*(*envIntl)(unsafe.Pointer(e))).ext.AskExternalData(int64(eid), int64(did), readSpan(span))
}

//export cGetExternalData
func cGetExternalData(e *C.env_t, eid C.int64_t, vid C.int64_t) C.Span {
	key := [2]int64{int64(eid), int64(vid)}
	env := (*(*envIntl)(unsafe.Pointer(e)))
	if _, ok := env.extData[key]; !ok {
		env.extData[key] = copySpan(env.ext.GetExternalData(int64(eid), int64(vid)))
	}
	return env.extData[key]
}
