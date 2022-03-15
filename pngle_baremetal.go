//go:build baremetal
// +build baremetal

package pngle

/*
#include <stdio.h>
int vfprintf() {
    return 0;
}
*/
import "C"
