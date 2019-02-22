// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 21:30 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"math"
	"strings"
)

// P(Category|Word) = (P(Word|Category) * P(Category))/P(Word)
// Input data is, at the same time a category mapping
func makeBayessianMatcher(input [][]string) func(a []string) (int, float64) {
	cat2String := func(data []string) string {
		return strings.Join(data, "")
	}

	// for calculation P(Category)
	category := map[string]uint32{}
	totalCategories := 0

	// for calculation P(Word|Category) P(word, given category)
	wordGivenCategory := map[string]map[string]uint32{}

	// iterate over data and do traning
	for _, row := range input {
		cat := cat2String(row)
		// training P(Category)
		category[cat] += 1
		totalCategories += 1

		wordGivenCategoryCat, ok := wordGivenCategory[cat]
		if !ok {
			wordGivenCategoryCat = map[string]uint32{}
		}

		for _, cell := range row {
			// training P(Word|Category) P(word, given category)
			wordGivenCategoryCat[cell] += 1
		}
		wordGivenCategory[cat] = wordGivenCategoryCat
	}

	// so, no zero in denominator
	if totalCategories == 0 {
		totalCategories = 1
	}

	// evaluate result of training
	return func(words []string) (int, float64) {
		bestCatIndex := 0
		bestP := 0.0
		for index, catSlice := range input {
			totalP := 0.0
			catString := cat2String(catSlice)

			// P(Category)
			categoryAppearTimes, ok := category[catString]
			if !ok {
				categoryAppearTimes = 1
			}

			pCat := math.Log(math.Log(float64(categoryAppearTimes)) - math.Log(float64(totalCategories)))
			totalP = pCat

			wordGivenCategory, ok := wordGivenCategory[catString]
			wordGivenCategoryLen := len(wordGivenCategory)
			if !ok {
				wordGivenCategory = map[string]uint32{}
				wordGivenCategoryLen = 1
			}

			for _, wordItem := range words {
				// calculate P(Word|Category), P(Word, given Category)
				wordGivenCategoryAppearsTimes, ok := wordGivenCategory[wordItem]
				if !ok {
					wordGivenCategoryAppearsTimes = 1
				}
				pWordGivenCategory := math.Log(float64(wordGivenCategoryAppearsTimes)) - math.Log(float64(wordGivenCategoryLen))
				totalP += pWordGivenCategory
			}
			if totalP > bestP {
				bestP = totalP
				bestCatIndex = index
			}
		}
		return bestCatIndex, bestP
	}
}
