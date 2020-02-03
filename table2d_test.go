package ascii_graph

import "testing"

func TestTable(t *testing.T) {
	table := Table2d{}

	row_p := []string{"name", "eeeage", "genderoaee"}
	col_p := []string{"0", "1"}

	data := [][]string{[]string{"bbaarry", "34", "w"}, []string{"oxlade", "13", "m"}}

	table.SetColumnProperty(col_p)
	table.SetRowProperty(row_p)
	table.SetData(data)

	table.Draw()
}
