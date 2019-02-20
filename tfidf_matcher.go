// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 00:19 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"math"
)

func makeTFIDFMatcher(input [][]string) func(a []string) ([]string, float64) {
	idf := map[string]uint32{}
	totalDocs := len(input)

	for _, row := range input {
		row2insert := map[string]bool{}
		for _, cell := range row {
			row2insert[cell] = true
		}
		for cell2insert, _ := range row2insert {
			idf[cell2insert] += 1
		}
	}

	return func(query []string) ([]string, float64) {
		bestResponse := []string{}
		bestResponseScore := 0.0

		for _, response := range input {
			termsFrequency := map[string]uint32{}
			maxTermFrequency := uint32(0)

			for _, word := range query {
				for _, response := range response {
					if response == word {
						termsFrequencyItem := termsFrequency[response]
						termsFrequencyItem += 1
						if termsFrequencyItem > maxTermFrequency {
							maxTermFrequency = termsFrequencyItem
						}
						termsFrequency[response] = termsFrequencyItem
					}
				}
			}
			result := float64(0)
			for term, termFrequency := range termsFrequency {
				tf := 0.5 + 0.5*float64(termFrequency/maxTermFrequency)
				idf := math.Log(float64(totalDocs) / (float64(1 + idf[term])))
				result += tf * idf
			}
			tfidf := result / float64(len(termsFrequency))
			if tfidf > bestResponseScore {
				bestResponse = response
				bestResponseScore = tfidf
			}
		}
		return bestResponse, bestResponseScore
	}

}
