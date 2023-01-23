//go:build !linux || !muslc

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
import "C"
