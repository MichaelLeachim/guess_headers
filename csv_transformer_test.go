// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 18:22 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTakeSeed(t *testing.T) {
	csv, err := ReadCSVFile("testdata/input1.csv", ',')
	assert.Equal(t, err, nil)
	assert.Equal(t, csv[0], "")
	transposedCsv := TransposeMatrix(csv)
	assert.Equal(t, transposedCsv[0], "")

}
