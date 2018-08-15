# go-table library 
Format output into a table

## usage
```
table := New(16, 4, 4, false)
for row := 0; row < 10; row ++ {
    var columns []string
    for col := 0; col < 10; col++ {
        columns = append(columns, fmt.Sprintf("%dX%d", row, col))
    }
    table.Row(columns...)
}
table.Print()
```