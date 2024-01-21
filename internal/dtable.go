package internal

import "fmt"

type DataType int

const (
	Int DataType = iota
	Float
	String
	Bool
)

type Column struct {
	Name     string
	DataType DataType
	Data     []interface{}
}

type DataTable struct {
	Columns []Column
}

// NewDataTable Create a new data table
func NewDataTable(columns ...Column) *DataTable {
	return &DataTable{Columns: columns}
}

// AddColumn Append a new column to the data table
func (dt *DataTable) AddColumn(column Column) {
	dt.Columns = append(dt.Columns, column)
}

// GetColumn Fetch column by index
func (dt *DataTable) GetColumn(index int) (Column, error) {
	if index < 0 || index >= len(dt.Columns) {
		return Column{}, fmt.Errorf("index out of bounds")
	}
	return dt.Columns[index], nil
}

// GetCell Fetch a specific cell's data
func (dt *DataTable) GetCell(colIndex int, rowIndex int) (interface{}, error) {
	column, err := dt.GetColumn(colIndex)
	if err != nil {
		return nil, err
	}
	if rowIndex < 0 || rowIndex >= len(column.Data) {
		return nil, fmt.Errorf("row index out of bounds")
	}
	return column.Data[rowIndex], nil
}
