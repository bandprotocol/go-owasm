package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "bindings.h"
//
// typedef Span (*get_calldata_fn)(env_t*);
// Span cGetCalldata_cgo(env_t *e);
// typedef void (*set_return_data_fn)(env_t*, Span);
// void cSetReturnData_cgo(env_t *e, Span data);
// typedef int64_t (*get_ask_count_fn)(env_t*);
// int64_t cGetAskCount_cgo(env_t *e);
// typedef int64_t (*get_min_count_fn)(env_t*);
// int64_t cGetMinCount_cgo(env_t *e);
// typedef int64_t (*get_ans_count_fn)(env_t*);
// int64_t cGetAnsCount_cgo(env_t *e);
import "C"
import "unsafe"

func Compile(code []byte) []byte {
	inputSpan := copySpan(code)
	defer freeSpan(inputSpan)
	outputSpan := newSpan(1 * 1024 * 1024)
	defer freeSpan(outputSpan)
	C.do_compile(inputSpan, &outputSpan)
	return readSpan(outputSpan)
}

func Run(code []byte, env EnvInterface) int32 {
	codeSpan := copySpan(code)
	defer freeSpan(codeSpan)
	envIntl := createEnvIntl(env)
	defer destroyEnvIntl(envIntl)
	C.do_run(codeSpan, C.Env{
		env: (*C.env_t)(unsafe.Pointer(envIntl)),
		dis: C.EnvDispatcher{
			get_calldata:    C.get_calldata_fn(C.cGetCalldata_cgo),
			set_return_data: C.set_return_data_fn(C.cSetReturnData_cgo),
			get_ask_count:   C.get_ask_count_fn(C.cGetAskCount_cgo),
			get_min_count:   C.get_min_count_fn(C.cGetMinCount_cgo),
			get_ans_count:   C.get_ans_count_fn(C.cGetAnsCount_cgo),
		},
	})
	return 10
}
