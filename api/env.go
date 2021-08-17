package api

// #include "bindings.h"
import "C"
import (
	"unsafe"
)

type EnvInterface interface {
	GetCalldata() []byte
	SetReturnData([]byte) error
	GetAskCount() uint64
	GetMinCount() uint64
	GetPrepareTime() int64
	GetExecuteTime() (int64, error)
	GetAnsCount() (uint64, error)
	AskExternalData(eid uint64, did uint64, data []byte) error
	GetExternalDataStatus(eid uint64, vid uint64) (int64, error)
	GetExternalData(eid uint64, vid uint64) ([]byte, error)
}

type envIntl struct {
	ext EnvInterface
}

func createEnvIntl(ext EnvInterface) *envIntl {
	return &envIntl{ext: ext}
}

//export cGetCalldata
func cGetCalldata(e *C.env_t, calldata *C.Span) C.Error {
	data := (*(*envIntl)(unsafe.Pointer(e))).ext.GetCalldata()
	return writeSpan(calldata, data)
}

//export cSetReturnData
func cSetReturnData(e *C.env_t, span C.Span) C.Error {
	err := (*(*envIntl)(unsafe.Pointer(e))).ext.SetReturnData(readSpan(span))
	if err != nil {
		return toCError(err)
	}
	return C.Error_NoError
}

//export cGetAskCount
func cGetAskCount(e *C.env_t) C.uint64_t {
	return C.uint64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetAskCount())
}

//export cGetMinCount
func cGetMinCount(e *C.env_t) C.uint64_t {
	return C.uint64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetMinCount())
}

//export cGetPrepareTime
func cGetPrepareTime(e *C.env_t) C.int64_t {
	return C.int64_t((*(*envIntl)(unsafe.Pointer(e))).ext.GetPrepareTime())
}

//export cGetExecuteTime
func cGetExecuteTime(e *C.env_t, val *C.int64_t) C.Error {
	v, err := (*(*envIntl)(unsafe.Pointer(e))).ext.GetExecuteTime()
	if err != nil {
		return toCError(err)
	}
	*val = C.int64_t(v)
	return C.Error_NoError
}

//export cGetAnsCount
func cGetAnsCount(e *C.env_t, val *C.uint64_t) C.Error {
	v, err := (*(*envIntl)(unsafe.Pointer(e))).ext.GetAnsCount()
	if err != nil {
		return toCError(err)
	}
	*val = C.uint64_t(v)
	return C.Error_NoError
}

//export cAskExternalData
func cAskExternalData(e *C.env_t, eid C.int64_t, did C.int64_t, span C.Span) C.Error {
	err := (*(*envIntl)(unsafe.Pointer(e))).ext.AskExternalData(uint64(eid), uint64(did), readSpan(span))
	if err != nil {
		return toCError(err)
	}
	return C.Error_NoError
}

//export cGetExternalDataStatus
func cGetExternalDataStatus(e *C.env_t, eid C.int64_t, vid C.int64_t, status *C.int64_t) C.Error {
	s, err := (*(*envIntl)(unsafe.Pointer(e))).ext.GetExternalDataStatus(uint64(eid), uint64(vid))
	if err != nil {
		return toCError(err)
	}
	*status = C.int64_t(s)
	return C.Error_NoError
}

//export cGetExternalData
func cGetExternalData(e *C.env_t, eid C.int64_t, vid C.int64_t, data *C.Span) C.Error {
	env := (*(*envIntl)(unsafe.Pointer(e)))
	extData, err := env.ext.GetExternalData(uint64(eid), uint64(vid))
	if err != nil {
		return toCError(err)
	}
	return writeSpan(data, extData)
}
