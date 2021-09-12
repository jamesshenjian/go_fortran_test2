package go_fortran_test

// #cgo LDFLAGS: -lgfortran
import "C"

func Hello() {
	C.hello()
}
