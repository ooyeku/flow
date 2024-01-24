package main

import (
	dt "flow/internal" // Use the correct import path for the "tables" package
	"fmt"
)

func main() {
	// Creating Columns
	col1 := dt.Column{
		Name:     "ID",
		DataType: dt.Int,
		Data:     []interface{}{1, 2, 3, 4, 5},
	}

	col2 := dt.Column{
		Name:     "Name",
		DataType: dt.String,
		Data:     []interface{}{"John", "Doe", "Jane", "Smith", "Bob"},
	}

	col3 := dt.Column{
		Name:     "Age",
		DataType: dt.Int,
		Data:     []interface{}{35, 28, 50, 23, 44},
	}

	// Creating DataTable
	dtable := dt.NewDataTable(col1, col2, col3)

	// Adding a column to DataTable
	col4 := dt.Column{
		Name:     "Salary",
		DataType: dt.Float,
		Data:     []interface{}{1000.5, 2000.25, 3000.75, 4000.0, 5000.85},
	}
	dtable.AddColumn(col4)

	// Fetching a cell's data
	cellData, err := dtable.GetCell(0, 1) // Get data from 1st row of 1st column (ID of 2nd employee)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data in the cell is: %v", cellData)
	}
}
