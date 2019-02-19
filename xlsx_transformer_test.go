// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 19:44 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadXLSXFile(t *testing.T) {
	testShouldEqual, _ := ReadCSVFile("testdata/input1.csv", ',')
	result, err := ReadXLSXFile("testdata/input1.xlsx", "input1")
	assert.Equal(t, err, nil)
	assert.Equal(t, testShouldEqual, result)
}

func TestWriteXLSXFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "tmp_xlsx_file")
	assert.Equal(t, err, nil)
	// remove after use
	defer os.Remove(tmpfile.Name())

	// read data for writing
	csv, _ := ReadCSVFile("testdata/input1.csv", ',')
	// write it into temp file
	assert.Equal(t, WriteXLSXFile(tmpfile.Name(), "sampleSheet", csv), nil)
	// read again
	xlsx, err := ReadXLSXFile(tmpfile.Name(), "sampleSheet")
	// assume equals
	assert.Equal(t, csv, xlsx)

}
