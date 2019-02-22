// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 18:15 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func WriteCSVFile(fpath string, data [][]string) error {
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func ReadCSVFile(fpath string, comma rune) ([][]string, error) {
	result := [][]string{}

	csvFile, err := os.Open(fpath)
	if err != nil {
		return result, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = comma
	for {
		line, err := reader.Read()
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return result, err
		}
		result = append(result, line)
	}
}
