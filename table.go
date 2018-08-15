package table

import (
    "fmt"
    "text/tabwriter"
    "os"
    "strconv"
    "strings"
    "regexp"
)

type Table struct {
    indentation int
    minWidth    int
    maxWidth    int
    writer      *tabwriter.Writer
    numberRows  bool
    rows        int
}

const flags = 0 //| tabwriter.Debug

func New(maxWidth, minWidth, indentation int, numberRows bool, header ...string) *Table {
    w := new(tabwriter.Writer)
    w.Init(os.Stdout, minWidth, 0, 4, ' ', flags)
    tw := &Table{
        indentation: indentation,
        minWidth:    minWidth,
        maxWidth:    maxWidth,
        rows:        0,
        numberRows:  numberRows,
        writer:      w}

    if len(header) > 0 {
        tw.Header(header...)
    }

    return tw
}

func (t *Table) split(value string) []string {
    space := regexp.MustCompile(`\s+`)
    value = space.ReplaceAllString(value, " ")

    var result []string
    var row string
    chars := []byte(value)
    for _, char := range chars {
        if row == "" && string(char) == " " {
            // rows will not start with whitespace
            continue
        }
        row = row + string(char)
        if len(row) >= t.maxWidth {
            whitespaceIndex := strings.LastIndex(row, " ")
            if whitespaceIndex > 0 && whitespaceIndex < len(row) {
                chars := []rune(row)
                left := string(chars[:whitespaceIndex])
                right := string(chars[whitespaceIndex+1:])
                result = append(result, left)
                row = right
            } else {
                result = append(result, row)
                row = ""
            }
        }
    }
    if row != "" {
        result = append(result, row)
    }
    return result
}

func (t *Table) Size() int {
    return t.rows
}

func (t *Table) row(columns ...string) error {
    if len(columns) == 0 {
        return nil
    }

    maxSplits := 0
    all := make([][]string, len(columns))
    for i, value := range columns {
        all[i] = t.split(value)
        if len(all[i]) > maxSplits {
            maxSplits = len(all[i])
        }
    }

    for j := 0; j < maxSplits; j++ {
        if t.indentation > 0 {
            if _, err := fmt.Fprintf(t.writer, "%s", strings.Repeat(" ", t.indentation)); err != nil {
                return err
            }
        }
        for i := 0; i < len(columns); i++ {
            var value string
            if j < len(all[i]) {
                value = all[i][j]
            }
            if _, err := fmt.Fprintf(t.writer, "%s\t", value); err != nil {
                return err
            }
        }
        if _, err := fmt.Fprintln(t.writer); err != nil {
            return err
        }
    }
    return nil
}

func (t *Table) Header(columns ...string) error {
    if t.numberRows {
        columns = append([]string{"#"}, columns...)
    }
    return t.row(columns...)
}

func (t *Table) Row(columns ...string) error {
    if t.numberRows {
        columns = append([]string{strconv.Itoa(t.rows + 1)}, columns...)
    }
    if err := t.row(columns...); err != nil {
        return err
    }
    t.rows++
    return nil
}

func (t *Table) Print() error {
    if err := t.writer.Flush(); err != nil {
        return err
    }
    return nil
}
