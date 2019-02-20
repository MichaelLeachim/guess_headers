// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 20:22 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	csv, _ := ReadCSVFile("testdata/input1.csv", ',')
	json, err := ReadJSONFile("testdata/input1.json")
	assert.Equal(t, err, nil)
	assert.Equal(t, csv, json)
}

func TestWriteJSONFile(t *testing.T) {

}
