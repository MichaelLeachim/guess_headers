// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 18:22 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	csv, err := ReadCSVFile("testdata/input1.csv", ',')
	assert.Equal(t, err, nil)
	assert.Equal(t, csv[0],
		[]string{"date", "docid", "title", "author", "pubinfo", "", "description", "condition", "price", "shipping", "Place"})
	headers, body := ChunkOffHeaders(TransposeMatrix(csv))
	assert.Equal(t, headers, []string([]string{"date", "docid", "title", "author", "pubinfo", "", "description", "condition", "price", "shipping", "Place", ""}))
	assert.Equal(t, body[11], []string{"", "", "", "", ""})

}

func TestWriteCSVFile(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	tmpfile, err := ioutil.TempFile("", "tmp_csv_file")
	assert.Equal(t, err, nil)
	// remove after use
	defer os.Remove(tmpfile.Name())

	// read data for writing
	csv, _ := ReadCSVFile("testdata/input1.csv", ',')
	// write it into temp file
	assert.Equal(t, WriteCSVFile(tmpfile.Name(), csv), nil)
	// read again
	csv2, err := ReadCSVFile(tmpfile.Name(), ',')
	log.Debug("Name of the temp file is: ", tmpfile.Name())

	assert.Equal(t, err, nil)
	// assume equals
	assert.Equal(t, csv, csv2)
}
