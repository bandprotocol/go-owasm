//go:build linux && muslc

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm_muslc -lgmp
import "C"
