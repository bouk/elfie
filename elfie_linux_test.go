package elfie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElfie(t *testing.T) {
	_, err := Elfie()
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
