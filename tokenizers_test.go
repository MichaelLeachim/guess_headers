// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 19:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeUnidecode(t *testing.T) {
	assert.Equal(t, TokenizeAlphaNumericOnly("Привет hello !№;%:? world*"), "")
}

func TestTokenizeAlphaNumericOnly(t *testing.T) {
	assert.Equal(t, TokenizeAlphaNumericOnly("Привет hello !№;%:? world*"), "hello world")
}

func TestTokenizeNumbers(t *testing.T) {
	assert.Equal(t, TokenizeNumbers("1920"), "1900 900 20 0")
	assert.Equal(t, TokenizeNumbers("0.00345"), "0.003 0.0004 0.00005")
}
