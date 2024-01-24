package internal

import (
	"errors"
	"fmt"
)

type dataType int

const (
	Int dataType = iota
	String
	Float
)

type column struct {
	name     string
	dataType dataType
	data     []interface{}
}

type dataTable struct {
	columns []column
}

func newColumn(name string, dataType dataType, data []interface{}) column {
	return column{
		name:     name,
		dataType: dataType,
		data:     data,
	}
}

func newDataTable(cols ...column) (*dataTable, error) {
	length := len(cols[0].data)

	for _, c := range cols {
		if len(c.data) != length {
			return nil, errors.New("all columns must have the same number of data elements")
		}
	}

	return &dataTable{columns: cols}, nil
}

func (dt *dataTable) addColumn(c column) error {
	if len(c.data) != len(dt.columns[0].data) {
		return errors.New("new column must have the same number of data elements as existing columns")
	}

	dt.columns = append(dt.columns, c)
	return nil
}

func (dt *dataTable) getCell(colIndex, rowIndex int) (interface{}, error) {
	if colIndex < 0 || colIndex >= len(dt.columns) {
		return nil, fmt.Errorf("column index %d out of range", colIndex)
	}

	if rowIndex < 0 || rowIndex >= len(dt.columns[colIndex].data) {
		return nil, fmt.Errorf("row index %d out of range for column %d", rowIndex, colIndex)
	}

	return dt.columns[colIndex].data[rowIndex], nil
}

func dtableDemo() {
	// Creating Columns using newColumn function
	IDColumn := newColumn("ID", Int, []interface{}{1, 2, 3, 4, 5})
	NameColumn := newColumn("Name", String, []interface{}{"John", "Doe", "Jane", "Smith", "Bob"})
	AgeColumn := newColumn("Age", Int, []interface{}{35, 28, 50, 23, 44})

	// Creating DataTable using newDataTable function
	dtable, err := newDataTable(IDColumn, NameColumn, AgeColumn)
	if err != nil {
		fmt.Println("Error creating DataTable:", err)
	}

	// Adding a column to DataTable
	SalaryColumn := newColumn("Salary", Float, []interface{}{1000.5, 2000.25, 3000.75, 4000.0, 5000.85})
	err = dtable.addColumn(SalaryColumn)
	if err != nil {
		fmt.Println("Error adding column:", err)
	}

	// Fetching a cell's data
	cellData, err := dtable.getCell(0, 1) // Get data from 1st row of 1st column (ID of 2nd employee)
	if err != nil {
		fmt.Println("Error fetching cell data:", err)
	}

	fmt.Printf("Data in the cell is: %v", cellData)
}
