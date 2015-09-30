package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	fmt.Println("Opening original.xlsx")
	originalXlFile, err := xlsx.OpenFile("original.xlsx")

	if err != nil {
		panic(err)
	}

	originalXlFile.Save("after_write.xlsx")

	fmt.Println("Opening after_write.xlsx")
	_, err = xlsx.OpenFile("after_write.xlsx")

	if err != nil {
		panic(err)
	}
}
