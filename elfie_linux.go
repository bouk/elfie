package elfie

import (
	"bytes"
	"debug/elf"
	"debug/gosym"
	"reflect"
	"unsafe"
)

var dma = *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
	Data: 0x0,
	Len:  0x7FFFFFFFF,
	Cap:  0x7FFFFFFFF,
}))

// Elfie parses the ELF file for the current process from memory
func Elfie() (*elf.File, error) {
	r := bytes.NewReader(dma[0x400000:])
	return elf.NewFile(r)
}

func table() (*gosym.Table, error) {
	elf, err := Elfie()
	if err != nil {
		return nil, err
	}

	var (
		textStart       uint64
		symtab, pclntab []byte
	)

	if sect := elf.Section(".text"); sect != nil {
		textStart = sect.Addr
	}

	if sect := elf.Section(".gosymtab"); sect != nil {
		if symtab, err = sect.Data(); err != nil {
			return nil, err
		}
	}

	if sect := elf.Section(".gopclntab"); sect != nil {
		if pclntab, err = sect.Data(); err != nil {
			return nil, err
		}
	}

	return gosym.NewTable(symtab, gosym.NewLineTable(pclntab, textStart))
}
