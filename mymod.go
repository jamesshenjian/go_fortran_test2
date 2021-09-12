package go_fortran_test2

// based on https://stackoverflow.com/a/53360659/1292467

// #cgo LDFLAGS: -lgfortran
// void hello();
import "C"

func Hello() {
	C.hello()
}
