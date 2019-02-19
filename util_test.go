// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 20:17 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransposeMatrix(t *testing.T) {
	matrix := [][]string{
		[]string{"H", "e", "l", "l", "o"},
		[]string{"W", "o", "r", "l", "d"},
	}
	assert.Equal(t, TransposeMatrix(TransposeMatrix(matrix)), matrix)
	assert.Equal(t, TransposeMatrix(matrix), "")
}

func TestTokenizeNumbers(t *testing.T) {
	assert.Equal(t, TokenizeNumbers("1920"), "1900 900 20 0")
	assert.Equal(t, TokenizeNumbers("0.00345"), "0.003 0.0004 0.00005")
}
