//go:build linux && !muslc && amd64

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm.x86_64 -lgmp
import "C"
