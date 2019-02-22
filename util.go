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
func newTestTriplet(score float64, left, right string, kind int) Triplet {
	return Triplet{Left: strings.Split(left, " "), Right: strings.Split(right, " "), Score: score, Kind: kind}
}

func RepeatString(repeat string, times int) []string {
	result := []string{}
	for i := 0; i <= times; i++ {
		result = append(result, repeat)
	}
	return result
}

// will transpose input matrix
func TransposeMatrix(input [][]string) [][]string {

	// instantiate output
	output := [][]string{}
	for i := 0; i <= len(input[0]); i++ {
		output = append(output, make([]string, len(input)))
	}

	// populate it with data
	for index4row, row := range input {
		for index4col, cell := range row {
			output[index4col][index4row] = cell
		}
	}
	return output
}
