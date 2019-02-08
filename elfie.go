// Package elfie takes an 'ELF-selfie' of the current process
package elfie // import "bou.ke/elfie"

import (
	"debug/gosym"
)

// Table constructs the symbol table for the current process
func Table() (*gosym.Table, error) {
	return table()
}
