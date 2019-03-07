package elfie

import (
	"bytes"
	"debug/elf"
	"debug/gosym"
)

// Elf parses the ELF file for the current process from memory
func Elf() (*elf.File, error) {
	r := bytes.NewReader(dma[0x400000:])
	return elf.NewFile(r)
}

func table() (*gosym.Table, error) {
	elf, err := Elf()
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
