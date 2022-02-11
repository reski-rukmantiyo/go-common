package tools

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx/v3"
)

func WriteToCell(sheet *xlsx.Sheet, rowX *xlsx.Row, col int, value string, params ...bool) {
	cell := rowX.AddCell()
	if value == "" {
		value = "0"
	}
	if params != nil && params[0] {
		idx := strings.Index(value, ".")
		if idx > 0 {
			value = value[0:idx]
		}
		val, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		cell.SetInt(val)
		cell.SetFormat("#")
	} else {
		cell.SetValue(value)
	}
}
