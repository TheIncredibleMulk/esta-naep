package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type estaTlv struct {
	tlvType                      string
	tlvSubType                   int
	name                         string
	stdSupported                 string
	status                       string
	pidStatus                    string
	suggestedThirdPartyTlv       string
	suggestedThirdPartyTlvSource string
	naepGuidance                 string
	description                  string
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
	proposedLenBytes             int
	proposedLenBits              int
	reference                    string
	featureClauseReference       string
	OptionalOrMandatory          string
	naepUse                      string
	naepReUse                    string
	naep                         bool
}

var rowToStruct = map[int]string{
	0: "tlvType",
	1: "tlvSubType",
	2: "name",
	3: "stdSupported",
	4: "status",
	5: "pidStatus",
	7: "suggestedThirdPartyTlv",
	8: "suggestedThirdPartyTlvSource",
	9: "naepGuidance",
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

	var tlvMap []estaTlv
	var tlvFieldRoster []estaTlv
	var tlvThirdParty []estaTlv

	// Get value from the cell by given worksheet name and cell reference.
	// cell, err := f.GetCellValue("ESTA TLV Map", "A1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.

	// ESTA TLV Map
	rows, err := f.GetRows("ESTA TLV Map")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		var tlv estaTlv
		if i > 0 {
			for j, colCell := range row {
				switch j {
				case 0:
					tlv.tlvType = colCell
				case 1:
					a, err := strconv.Atoi(colCell)
					if err != nil {
						fmt.Println("Error string conversation to int")
					}
					tlv.tlvSubType = a
				case 2:
					tlv.name = colCell
				case 3:
					tlv.description = colCell
				case 4:
					tlv.stdSupported = colCell
				case 5:
					tlv.status = colCell
				case 6:
					tlv.pidStatus = colCell
				case 7:
					a, err := strconv.Atoi(colCell)
					if err != nil {
						fmt.Println("Error string conversation to int")
					}
					tlv.proposedLenBytes = a
				case 8:
					a, err := strconv.Atoi(colCell)
					if err != nil {
						fmt.Println("Error string conversation to int")
					}
					tlv.proposedLenBits = a
				case 9:
					tlv.estaTlvName = colCell
				case 10:
					if colCell == "Yes" || colCell == "Y" || colCell == "yes" || colCell == "y" {
						tlv.varibleLength = true
					}
					if colCell == "No" || colCell == "N" || colCell == "no" || colCell == "n" {
						tlv.varibleLength = false
					}
				case 11:
					tlv.sender = colCell
				case 12:
					tlv.msByte = colCell
				case 13:
					tlv.msBit = colCell
				case 14:
					tlv.lsByte = colCell
				case 15:
					tlv.lsBit = colCell
				case 16:
					tlv.fieldDescription = colCell
				case 17:
					tlv.comments = colCell
				case 18:
					tlv.suggestedThirdPartyTlv = colCell
				case 19:
					tlv.suggestedThirdPartyTlvSource = colCell
				case 20:
					tlv.naepGuidance = colCell
				}
				tlvMap = append(tlvMap, tlv)
			}
		}
		rows, err = f.GetRows("ESTA TLV Field Roster")
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, row := range rows {
			var tlv estaTlv
			if i > 0 {
				for j, colCell := range row {
					switch j {
					case 0:
						tlv.tlvType = colCell
					case 1:
						a, err := strconv.Atoi(colCell)
						if err != nil {
							fmt.Println("Error string conversation to int")
						}
						tlv.tlvSubType = a
					case 2:
						tlv.name = colCell
					case 3:
						tlv.reference = colCell
					case 4:
						tlv.featureClauseReference = colCell
					case 5:
						tlv.OptionalOrMandatory = colCell
					case 7:
						tlv.naepUse = colCell
					case 8:
						tlv.naepReUse = colCell
					case 9:
						tlv.naepGuidance = colCell
					}
				}
				tlvFieldRoster = append(tlvFieldRoster, tlv)
			}
		}
	}

	// 3rd Party TLV Roster
	rows, err = f.GetRows("3rd Party TLV Roster")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		var tlv estaTlv
		if i > 0 {
			for j, colCell := range row {
				switch j {
				case 0:
					tlv.tlvType = colCell
				case 1:
					a, err := strconv.Atoi(colCell)
					if err != nil {
						fmt.Println("Error string conversation to int")
					}
					tlv.tlvSubType = a
				case 2:
					tlv.name = colCell
				case 3:
					tlv.reference = colCell
				case 4:
					tlv.status = colCell
				case 5:
					tlv.pidStatus = colCell
				case 7:
					tlv.suggestedThirdPartyTlv = colCell
				case 8:
					tlv.suggestedThirdPartyTlvSource = colCell
				case 9:
					tlv.naepGuidance = colCell
				}
			}
			tlvThirdParty = append(tlvThirdParty, tlv)
		}
	}
	fmt.Println("tlv Map: ")
	fmt.Println(tlvMap)
	fmt.Println()
	fmt.Println("tlv Field Roster: ")
	fmt.Println(tlvFieldRoster)
	fmt.Println()
	fmt.Println("3rd Party TLV Roster: ")
	fmt.Println(tlvThirdParty)

}
