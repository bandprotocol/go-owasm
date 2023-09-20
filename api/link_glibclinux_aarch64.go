//go:build linux && !muslc && arm64

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm.aarch64 -lgmp
import "C"
