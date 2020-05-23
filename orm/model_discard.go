package orm

import (
	"github.com/nibeh/pg/v9/types"
)

type Discard struct {
	hookStubs
}

var _ Model = (*Discard)(nil)

func (Discard) Init() error {
	return nil
}

func (m Discard) NextColumnScanner() ColumnScanner {
	return m
}

func (m Discard) AddColumnScanner(ColumnScanner) error {
	return nil
}

func (m Discard) ScanColumn(colIdx int, colName string, rd types.Reader, n int) error {
	return nil
}
