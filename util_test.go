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

	// how it works on empty input?
	assert.Equal(t, [][]string{}, "")

	// how it works on sparse input?
	assert.Equal(t, [][]string{[]string{"H"}, []string{"H", "A"}}, "")

}
