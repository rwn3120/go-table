package table

import (
    "testing"
    "fmt"
)

func TestTable(t *testing.T) {
    fmt.Println("demo1:", t.Name())
    table := New(16, 4, 4, false)
    for row := 0; row < 10; row ++ {
        var columns []string
        for col := 0; col < 10; col++ {
            columns = append(columns, fmt.Sprintf("%dX%d", row, col))
        }
        table.Row(columns...)
    }
    table.Print()
}

func TestTableWithRowNumbers(t *testing.T) {
    fmt.Println("demo2:", t.Name())
    table := New(16, 4, 4, true)
    for row := 0; row < 10; row ++ {
        var columns []string
        for col := 0; col < 10; col++ {
            columns = append(columns, fmt.Sprintf("%dX%d", row, col))
        }
        table.Row(columns...)
    }
    table.Print()
}


func TestTableWithHeader(t *testing.T) {
    fmt.Println("demo3:", t.Name())

    var headers []string
    for col := 0; col < 10; col++ {
        headers = append(headers, fmt.Sprintf("header %d", col))
    }

    table := New(16, 4, 4, false, headers...)
    for row := 0; row < 10; row ++ {
        var columns []string
        for col := 0; col < 10; col++ {
            columns = append(columns, fmt.Sprintf("%dX%d", row, col))
        }
        table.Row(columns...)
    }
    table.Print()
}
func TestTableWithHeaderWithRowNumbers(t *testing.T) {
    fmt.Println("demo4:", t.Name())

    var headers []string
    for col := 0; col < 10; col++ {
        headers = append(headers, fmt.Sprintf("header %d", col))
    }

    table := New(16, 4, 4, true, headers...)
    for row := 0; row < 10; row ++ {
        var columns []string
        for col := 0; col < 10; col++ {
            columns = append(columns, fmt.Sprintf("%dX%d", row, col))
        }
        table.Row(columns...)
    }
    table.Print()
}