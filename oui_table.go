package mactracker

import (
	_ "embed"
	"fmt"
	"strings"
)

// ouiTableData is a string so the embedded blob lives in read-only memory.
//
//go:embed oui_table.bin.gz
var ouiTableData string

// ouiTableErr holds the decode error for the embedded table, if any.
var ouiTableErr error

// OUITable contains IEEE registrations, lazily loaded from embedded binary data.
var OUITable = OuiDB{
	loadFunc: func() map[string]*OuiBlock {
		var blocks map[string]*OuiBlock
		blocks, ouiTableErr = decodeEmbeddedOUITable(ouiTableData)
		return blocks
	},
}

// decodeEmbeddedOUITable returns an empty table alongside any decode error.
func decodeEmbeddedOUITable(data string) (map[string]*OuiBlock, error) {
	blocks, err := decodeOUIDB(strings.NewReader(data))
	if err != nil {
		return map[string]*OuiBlock{}, fmt.Errorf("mactracker: decode embedded OUI table: %w", err)
	}
	return blocks, nil
}

// OUITableError loads the embedded IEEE table and reports any decode error.
// A table that fails to decode yields no IEEE matches instead of panicking.
func OUITableError() error {
	OUITable.load()
	return ouiTableErr
}
