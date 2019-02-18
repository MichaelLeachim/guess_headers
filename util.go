// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:46 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"strings"
)

// usually, there is no need in creating triplets by hand,
// but for test reasons
func newTestTriplet(score float64, left, right string) Triplet {
	return Triplet{Left: strings.Split(left, " "), Right: strings.Split(right, " "), Score: score}
}

// from here: https://github.com/git-time-metric/gtm/blob/master/util/string.go#L53-L88
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// [Tokenize] Tokenize numbers via splitting, i.e. 1924 => (1000 900 20 4)
func TokenizeNumbers(number string) string {
	for index, _ := range number {
		number[:index]

	}
	number[0]

}
