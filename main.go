package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type estaTlvMap struct {
	tlvType                      string
	tlvSubType                   int
	name                         string
	description                  string
	stdSupported                 string
	status                       string
	pidStatus                    string
	byteLength                   int
	bitLength                    int
	estaTlvName                  string
	varibleLength                bool
	sender                       string
	msByte                       string
	msBit                        string
	lsByte                       string
	lsBit                        string
	fieldDescription             string
	comments                     string
	suggestedThirdPartyTlv       string
	suggestedThirdPartyTlvSource string
	naepGuidance                 string
}

func main() {
	f, err := excelize.OpenFile("NAEP_TLV_MAPPING.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get value from the cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("ESTA TLV Map", "A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("ESTA TLV Field Roster")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
