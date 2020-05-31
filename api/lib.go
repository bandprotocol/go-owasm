package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "bindings.h"
//
// typedef int64_t (*get_ask_count_fn)(env_t*);
// int64_t cGetAskCount_cgo(env_t *e);
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

func Run(code []byte, env Env) int32 {
	codeSpan := copySpan(code)
	defer freeSpan(codeSpan)
	return int32(C.do_run(codeSpan, C.Env{
		env: (*C.env_t)(unsafe.Pointer(&env)),
		dis: C.EnvDispatcher{
			get_ask_count: C.get_ask_count_fn(C.cGetAskCount_cgo),
		},
	}))
}
