// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:09 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"strings"
)

// Guessing algorithm (ideas):
// [Reduce space] Take random N(100) fields of each column of each file (reducing step)
// [Tokenize] Tokenize numbers via splitting, i.e. 1924 => (1000 900 20 4)
// [Match] Match columns: dot product with each other, keep only the best match
// [Cleanup] Assume that there is only single one best match.
// If it is taken, there is nothing left (the match is nil)

func MatchAlgoHamming(input1, input2 []string) float64 {

}

// Chooses something from input2 that matches input1 in the best way
func CalculateBestMatch(input1 []string, input2 [][]string) Triplet {

}

// [Cleanup] Assume that there is only single best match. If it is taken, there is nothing left (the match is nil)
func CleanUp(data []Triplet) []Triplet {
	scores := map[string]float64{}
	for index, item := range data {
		score, ok := scores[strings.Join(item.Right, "")]
		if ok && score <= item.Score {
			scores[strings.Join(item.Right, "")] = item.Score
		}
	}
	for index, item := range data {
		score, ok := scores[strings.Join(item.Right, "")]
		if ok && score <= item.Score {
			data[index].Right = nil
			data[index].Score = 0
		}
	}
	return data
}

// returns match score
func MatchBetween(input1, input2 []string) float64 {

}
