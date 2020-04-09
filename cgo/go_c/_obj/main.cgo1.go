// Created by cgo - DO NOT EDIT

//line main.go:1
package main

//使用#cgo定义库路径

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lhello
// #include "hello.h"
import _ "unsafe"

func main() {
	(_Cfunc_hello)()
}
