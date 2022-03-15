package pngle

/*
//#cgo CFLAGS: -DPNGLE_DEBUG

#include "miniz.h"
#include "pngle.h"
int decodeFromBytes(unsigned char *b, int length, int scale);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func DecodeFromBytes(b []byte) {
	ret := C.decodeFromBytes((*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)), C.int(0))

	fmt.Printf("ret : %d\n", ret)
}

var pngleBuf [256]uint16

//export callbackFromPngle
func callbackFromPngle(x, y, w, h uint16, r, g, b, a uint8) {
	callback(x, y, w, h, r, g, b, a)
}

var callback func(x, y, w, h uint16, r, g, b, a uint8)

func SetCallback(fn func(x, y, w, h uint16, r, g, b, a uint8)) {
	callback = fn
}
