package ascii_graph

import "fmt"

type Table2d struct {
	data            [][]string
	col_property    []string
	row_property    []string
	largest_col     int
	largest_element []int
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

func (t *Table2d) DrawLine() {

	for i := 0; i < t.largest_col; i++ {
		fmt.Printf("-")
	}

	for i := 0; i < len(t.row_property); i++ {
		for j := 0; j < t.largest_element[i]+1; j++ {
			fmt.Printf("-")
		}
	}

	fmt.Printf("----\n")

}

func (t *Table2d) DrawDataRow() {

	for i, col_p := range t.col_property {

		fmt.Printf("|")

		fmt.Printf("%s ", col_p)
		for i := 0; i < t.largest_col-len(col_p); i++ {
			fmt.Printf(" ")
		}

		fmt.Printf("| ")

		for j, d := range t.data[i] {
			fmt.Printf("%s ", d)
			for k := 0; k < t.largest_element[j]-len(d); k++ {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}

}

func (t *Table2d) DrawRowProperty() {

	fmt.Printf("|")

	for i := 0; i < t.largest_col+1; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("| ")

	for j, row_p := range t.row_property {
		fmt.Printf("%s ", row_p)
		for i := 0; i < t.largest_element[j]-len(row_p); i++ {
			fmt.Printf(" ")

		}
	}

	fmt.Printf("|")

	fmt.Printf("\n")

}

func (t *Table2d) SetLargest() {
	data_col_length := len(t.data)
	if data_col_length != len(t.col_property) {
		err := fmt.Errorf("data length %d is not equlal to col length %d", data_col_length, len(t.col_property))
		panic(err.Error())
	}

	data_row_length := len(t.data[0])
	if data_row_length != len(t.row_property) {
		err := fmt.Errorf("data length %d is not equlal to col length %d", data_row_length, len(t.row_property))
		panic(err.Error())
	}

	t.largest_element = make([]int, len(t.row_property))

	t.largest_col = 0
	for _, col_p := range t.col_property {
		if t.largest_col < len(col_p) {
			t.largest_col = len(col_p)
		}
	}

	for i := 0; i < len(t.row_property); i++ {
		t.largest_element[i] = len(t.row_property[i])
		for j := 0; j < len(t.col_property); j++ {
			if t.largest_element[i] < len(t.data[j][i]) {
				t.largest_element[i] = len(t.data[j][i])
			}
		}

	}

}

func (t *Table2d) Draw() {
	t.SetLargest()
	t.DrawLine()
	t.DrawRowProperty()
	t.DrawLine()
	t.DrawDataRow()
	t.DrawLine()

}
