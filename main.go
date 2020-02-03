package main

import "fmt"

type Table2d struct {
	data         [][]string
	col_property []string
	row_property []string
}

func (t *Table2d) SetRowProperty(row []string) {
	t.row_property = row
}

func (t *Table2d) SetColumnProperty(col []string) {
	t.col_property = col
}

func (t *Table2d) SetData(data [][]string) {
	t.data = data
}

func (t *Table2d) Draw() {
	data_col_length := len(t.data)
	if data_col_length != len(t.col_property) {

	}

	data_row_length := len(t.data[0])
	if data_row_length != len(t.row_property) {

	}

	largest_element := make([]int, len(t.row_property))

	largest_col := 0
	for _, col_p := range t.col_property {
		if largest_col < len(col_p) {
			largest_col = len(col_p)
		}
	}

	for i := 0; i < len(t.row_property); i++ {
		largest_element[i] = len(t.row_property[i])
		for j := 0; j < len(t.col_property); j++ {
			if largest_element[i] < len(t.data[j][i]) {
				largest_element[i] = len(t.data[j][i])
			}
		}

	}

	for i := 0; i < largest_col; i++ {
		fmt.Printf("-")
	}

	for i := 0; i < len(t.row_property); i++ {
		for j := 0; j < largest_element[i]+1; j++ {
			fmt.Printf("-")
		}
	}
	fmt.Printf("--\n")

	fmt.Printf("|")

	for i := 0; i < largest_col+1; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("| ")

	for j, row_p := range t.row_property {
		fmt.Printf("%s ", row_p)
		for i := 0; i < largest_element[j]-len(row_p); i++ {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")

	for i := 0; i < largest_col; i++ {
		fmt.Printf("-")
	}

	for i := 0; i < len(t.row_property); i++ {
		for j := 0; j < largest_element[i]+1; j++ {
			fmt.Printf("-")
		}
	}

	fmt.Printf("--\n")

	for i, col_p := range t.col_property {

		fmt.Printf("|")

		fmt.Printf("%s ", col_p)
		for i := 0; i < largest_col-len(col_p); i++ {
			fmt.Printf(" ")
		}

		fmt.Printf("| ")

		for j, d := range t.data[i] {
			fmt.Printf("%s ", d)
			for k := 0; k < largest_element[j]-len(d); k++ {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

	for i := 0; i < largest_col; i++ {
		fmt.Printf("-")
	}

	for i := 0; i < len(t.row_property); i++ {
		for j := 0; j < largest_element[i]+1; j++ {
			fmt.Printf("-")
		}
	}

	fmt.Printf("--\n")

}

func main() {
	table := Table2d{}

	row_p := []string{"name", "eeeage", "genderoaee"}
	col_p := []string{"0", "1"}

	data := [][]string{[]string{"bbaarry", "34", "w"}, []string{"oxlade", "13", "m"}}

	table.SetColumnProperty(col_p)
	table.SetRowProperty(row_p)
	table.SetData(data)

	table.Draw()

}
