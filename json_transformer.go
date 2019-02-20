// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 20:10 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJSONFile(fpath string, data [][]string) error {
	head := data[0]
	body := data[1:]

	jsondata := []map[string]string{}
	for _, bodyItem := range body {
		row := map[string]string{}
		for index, cell := range bodyItem {
			nameOfCell := head[index]
			row[nameOfCell] = cell
		}
		jsondata = append(jsondata, row)
	}
	jsonAsBytes, err := json.Marshal(jsondata)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fpath, jsonAsBytes, 0644)
}

func ReadJSONFile(fpath string, comma rune) ([][]string, error) {
	pass
}
