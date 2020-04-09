package cgo

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

func Sqrt(p float32) (float32, error) {
    n, err := C.sqrt(C.double(p))
    return float32(n), err
}
