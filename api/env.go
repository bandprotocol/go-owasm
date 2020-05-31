package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "bindings.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type EnvInterface interface {
	GetCalldata() []byte
	GetAskCount() int64
	GetMinCount() int64
	GetAnsCount() int64
	// GetPrepareBlockTime() int64
	// GetAggregateBlockTime() int64
	// GetValidatorAddress(validatorIndex int64) ([]byte, error)
	// GetMaxResultSize() int64
	// GetMaxRawRequestDataSize() int64
	// RequestExternalData(dataSourceID int64, externalID int64, calldata []byte) error
	// GetExternalData(externalID int64, validatorIndex int64) ([]byte, uint32, error)
}

type envIntl struct {
	ext      EnvInterface
	calldata C.Span
}

func createEnvIntl(ext EnvInterface) *envIntl {
	return &envIntl{
		ext:      ext,
		calldata: copySpan(ext.GetCalldata()),
	}
}

func destroyEnvIntl(e *envIntl) {
	freeSpan(e.calldata)
}

//export cGetCalldata
func cGetCalldata(e *C.env_t) C.Span {
	return (*(*envIntl)(unsafe.Pointer(e))).calldata
}

//export cSetReturnData
func cSetReturnData(e *C.env_t, span C.Span) {
	data := readSpan(span)
	fmt.Println(string(data))
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
