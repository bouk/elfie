package elfie

import (
	"bytes"
	"debug/macho"
)

func Macho() (*macho.File, error) {
	r := bytes.NewReader(dma[0x1000000:])
	return macho.NewFile(r)
}
