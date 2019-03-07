// Package elfie takes an 'ELF-selfie' of the current process
package elfie // import "bou.ke/elfie"

import (
	"debug/gosym"
	"reflect"
	"unsafe"
)

var dma = *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
	Data: 0x0,
	Len:  0x7FFFFFFFF,
	Cap:  0x7FFFFFFFF,
}))

// Table constructs the symbol table for the current process
func Table() (*gosym.Table, error) {
	return table()
}
