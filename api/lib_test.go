package api

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

const TESTING_MEMORY_LIMIT = 32 * 1024 // KiB

func newTestVM(t *testing.T) (*Vm, func()) {
	vm, err := NewVm(TESTING_MEMORY_LIMIT)
	require.NoError(t, err)

	cleanup := func() {
		ReleaseCache(vm.cache)
	}
	return vm, cleanup
}

func readWatFile(fileName string) []byte {
	code, err := ioutil.ReadFile(fmt.Sprintf("./../wasm/%s.wat", fileName))
	if err != nil {
		panic(err)
	}
	return code
}

func readWasmFile(fileName string) []byte {
	code, err := ioutil.ReadFile(fmt.Sprintf("./../wasm/%s.wasm", fileName))
	if err != nil {
		panic(err)
	}
	return code
}

func TestSuccessWatToOwasm(t *testing.T) {
	code := readWatFile("test")
	wasm := wat2wasm(code)
	expectedWasm := readWasmFile("test")
	require.Equal(t, expectedWasm, wasm)
}

func TestFailCompileInvalidContent(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	code := []byte("invalid content")
	spanSize := 1 * 1024 * 1024
	_, err := vm.Compile(code, spanSize)
	require.Equal(t, ErrValidation, err)
}
func TestRuntimeError(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  i32.const 0
		  i32.const 0
		  i32.div_s
		  drop
		)
		(func)
		(memory 17)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

		`))
	code, _ := vm.Compile(wasm, spanSize)
	_, err := vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrRuntime, err)
}

func TestInvaildSignature(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(func (param i64 i64 i32 i64)
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(memory 17)
		(export "prepare" (func 0))
		(export "execute" (func 1)))
	  `))
	code, _ := vm.Compile(wasm, spanSize)
	_, err := vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrBadEntrySignature, err)
}

func TestGasLimit(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(memory 17)
		(export "prepare" (func 0))
		(export "execute" (func 1)))
	  `))
	code, err := vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	output, err := vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.NoError(t, err)
	require.Equal(t, RunOutput{GasUsed: 200032500000}, output)
	_, err = vm.Prepare(code, 175000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrOutOfGas, err)
}

func TestCompileErrorNoMemory(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	code, err := vm.Compile(wasm, spanSize)
	require.Equal(t, ErrBadMemorySection, err)
	require.Equal(t, []uint8([]byte{}), code)
}

func TestCompileErrorMinimumMemoryExceed(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(memory 512)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	_, err := vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	wasm = wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(memory 513)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	_, err = vm.Compile(wasm, spanSize)
	require.Equal(t, ErrBadMemorySection, err)
}

func TestCompileErrorSetMaximumMemory(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(func
		  (local $idx i32)
		  (local.set $idx (i32.const 0))
		  (block
			  (loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 10000)))
			  )
			))
		(func)
		(memory 17 20)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	code, err := vm.Compile(wasm, spanSize)
	require.Equal(t, ErrBadMemorySection, err)
	require.Equal(t, []uint8([]byte{}), code)
}

func TestCompileErrorCheckWasmImports(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(import "env" "beeb" (func (type 0)))
		(func
		(local $idx i32)
		(local.set $idx (i32.const 0))
		(block
				(loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 1000000000)))
				)
			)
		)
		(func)
		(memory 17)
		(data (i32.const 1048576) "beeb")
		(export "prepare" (func 0))
		(export "execute" (func 1)))
		`))
	code, err := vm.Compile(wasm, spanSize)
	require.Equal(t, ErrInvalidImports, err)
	require.Equal(t, []uint8([]byte{}), code)
}

func TestCompileErrorCheckWasmExports(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (func (param i64 i64 i32 i64) (result i64)))
		(import "env" "ask_external_data" (func (type 0)))
		(func
		(local $idx i32)
		(local.set $idx (i32.const 0))
		(block
				(loop
				(local.set $idx (local.get $idx) (i32.const 1) (i32.add) )
				(br_if 0 (i32.lt_u (local.get $idx) (i32.const 1000000000)))
				)
			)
		)
		(memory 17)
		(data (i32.const 1048576) "beeb")
		(export "prepare" (func 0)))
		`))
	code, err := vm.Compile(wasm, spanSize)
	require.Equal(t, ErrInvalidExports, err)
	require.Equal(t, []uint8([]byte{}), code)
}

func TestStackOverflow(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(func call 0)
		(func)
		(memory 10)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	code, _ := vm.Compile(wasm, spanSize)
	_, err := vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrRuntime, err)
}

func TestMemoryGrow(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(func
	i32.const 0
    (memory.grow (i32.const 1))
    i32.gt_s
	if
    	unreachable
    end
     )
		(func)
		(memory 10)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	code, _ := vm.Compile(wasm, spanSize)
	_, err := vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.NoError(t, err)

	wasm = wat2wasm([]byte(`(module
		(func
	i32.const 0
    (memory.grow (i32.const 1))
    i32.gt_s
	if
    	unreachable
    end
     )
		(func)
		(memory 512)
		(export "prepare" (func 0))
		(export "execute" (func 1)))

	  `))
	code, _ = vm.Compile(wasm, spanSize)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrRuntime, err)
}

func TestBadPointer(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (;0;) (func (param i64 i64)))
		(type (;1;) (func))
		(import "env" "set_return_data" (func (;0;) (type 0)))
		(func (type 1)
			i64.const 100000000
			i64.const 1
			call 0
			)
		(func)
		(memory 17)
		(export "prepare" (func 1))
		(export "execute" (func 2)))

		`))
	code, err := vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrMemoryOutOfBound, err)

	wasm = wat2wasm([]byte(`(module
		(type (;0;) (func (param i64 i64 i64 i64)))
		(type (;1;) (func))
		(import "env" "ask_external_data" (func (;0;) (type 0)))
		(func (type 1)
			i64.const 1
			i64.const 1
			i64.const 100000000
			i64.const 1
			call 0
			)
		(func)
		(memory 17)
		(export "prepare" (func 1))
		(export "execute" (func 2)))

		`))
	code, err = vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrMemoryOutOfBound, err)
}

func TestSpanTooSmall(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (;0;) (func (param i64 i64 i64 i64)))
		(type (;1;) (func))
		(import "env" "ask_external_data" (func (;0;) (type 0)))
		(func (type 1)
			i64.const 1
			i64.const 1
			i64.const 1
			i64.const 1024
			call 0
			)
		(func)
		(memory 17)
		(export "memory" (memory 0))
		(export "prepare" (func 1))
		(export "execute" (func 2)))
		`))
	code, err := vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.NoError(t, err)

	wasm = wat2wasm([]byte(`(module
		(type (;0;) (func (param i64 i64 i64 i64)))
		(type (;1;) (func))
		(import "env" "ask_external_data" (func (;0;) (type 0)))
		(func (type 1)
			i64.const 1
			i64.const 1
			i64.const 1
			i64.const 1025
			call 0
			)
		(func)
		(memory 17)
		(export "memory" (memory 0))
		(export "prepare" (func 1))
		(export "execute" (func 2)))
		`))
	code, err = vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrSpanTooSmall, err)
}

func TestBadImportSignature(t *testing.T) {
	vm, release := newTestVM(t)
	defer release()

	spanSize := 1 * 1024 * 1024
	wasm := wat2wasm([]byte(`(module
		(type (;0;) (func))
		(type (;1;) (func))
		(import "env" "set_return_data" (func (;0;) (type 0)))
		(func
			call 0)
		(func)
		(memory 17)
		(export "memory" (memory 0))
		(export "prepare" (func 1))
		(export "execute" (func 2)))

		`))
	code, err := vm.Compile(wasm, spanSize)
	require.NoError(t, err)
	_, err = vm.Prepare(code, 250000000000, 1024, NewMockEnv([]byte("")))
	require.Equal(t, ErrInstantiation, err)
}
