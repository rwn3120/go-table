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
output:
```
    0X0    0X1    0X2    0X3    0X4    0X5    0X6    0X7    0X8    0X9    
    1X0    1X1    1X2    1X3    1X4    1X5    1X6    1X7    1X8    1X9    
    2X0    2X1    2X2    2X3    2X4    2X5    2X6    2X7    2X8    2X9    
    3X0    3X1    3X2    3X3    3X4    3X5    3X6    3X7    3X8    3X9    
    4X0    4X1    4X2    4X3    4X4    4X5    4X6    4X7    4X8    4X9    
    5X0    5X1    5X2    5X3    5X4    5X5    5X6    5X7    5X8    5X9    
    6X0    6X1    6X2    6X3    6X4    6X5    6X6    6X7    6X8    6X9    
    7X0    7X1    7X2    7X3    7X4    7X5    7X6    7X7    7X8    7X9    
    8X0    8X1    8X2    8X3    8X4    8X5    8X6    8X7    8X8    8X9    
    9X0    9X1    9X2    9X3    9X4    9X5    9X6    9X7    9X8    9X9  
```