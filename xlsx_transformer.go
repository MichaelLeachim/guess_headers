// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 19:31 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"errors"
	"github.com/tealeg/xlsx"
	"strings"
)

// !!! will work only on one sheet of the file !!!
func WriteXLSXFile(fpath string, sheetName string, data [][]string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}

	for _, row := range data {
		row4file := sheet.AddRow()
		for _, cell := range row {
			cell4file := row4file.AddCell()
			cell4file.Value = cell
		}
	}
	return file.Save(fpath)
}

// !!! will work only on one sheet of the file !!!
func ReadXLSXFile(fpath string, sheetName string) ([][]string, error) {
	result := [][]string{}
	xlFile, err := xlsx.OpenFile(fpath)
	if err != nil {
		return result, err
	}
	for _, sheet := range xlFile.Sheets {
		if strings.TrimSpace(sheet.Name) != strings.TrimSpace(sheetName) {
			continue
		}
		for _, row := range sheet.Rows {
			row4result := []string{}
			for _, cell := range row.Cells {
				row4result = append(row4result, cell.String())
			}
			result = append(result, row4result)
		}
		return result, nil
	}
	return result, errors.New("File: " + fpath + "does not contain sheet: " + sheetName)
}
