package elfie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElf(t *testing.T) {
	_, err := Elf()
	require.NoError(t, err)
}

func TestTable(t *testing.T) {
	table, err := Table()
	require.NoError(t, err)
	for _, f := range table.Funcs {
		fmt.Println(f.Name)
	}
	t.Fatal()
}
