package pngle

/*
//#cgo CFLAGS: -DPNGLE_DEBUG

#include "miniz.h"
#include "pngle.h"
int pngle_decodeFromBytes(unsigned char *b, int length, int scale);
*/
import "C"
import (
	"unsafe"
)

func DecodeFromBytes(b []byte) {
	C.pngle_decodeFromBytes((*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)), C.int(0))
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
