// package elfie takes an 'ELF-selfie' of the current process
package elfie

import (
	"bytes"
	"debug/elf"
	"os"
	"reflect"
	"unsafe"

	"bou.ke/procmaps"
)

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func Elfie() (*elf.File, error) {
	mapping, err := procmaps.ReadSelf()
	if err != nil {
		return nil, err
	}
	exec, err := os.Executable()
	if err != nil {
		return nil, err
	}

	start := uintptr(0)
	end := uintptr(0)
	for _, m := range mapping {
		if m.Path != exec {
			break
		}
		if start == 0 {
			start = m.Start
		}
		end = m.End
	}
	mem := rawMemoryAccess(start, int(end-start))
	r := bytes.NewReader(mem)
	return elf.NewFile(r)
}
