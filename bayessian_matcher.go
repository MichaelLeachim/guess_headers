// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 21:30 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	log "github.com/sirupsen/logrus"
	"math"
	"strings"
)

// P(Category|Word) = (P(Word|Category) * P(Category))/P(Word)
// Input data is, at the same time a category mapping
func makeBayessianMatcher(input [][]string) func(a []string) (int, float64) {
	laplase_nom := 0.0000000001

	cat2String := func(data []string) string {
		return strings.Join(data, "")
	}

	// for calculation P(Category)
	category := map[string]uint32{}
	totalCategories := 0

	// for calculation P(Word|Category) P(word, given category)
	wordGivenCategory := map[string]map[string]uint32{}
	log.Debug("Starting training(iterating over input)")

	// iterate over data and do traning
	for _, row := range input {
		log.Debug("Row is: ", row)
		cat := cat2String(row)
		log.Debug("Stringified like: ", cat)
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
		log.Debug("WordGivenCategory (for category)", wordGivenCategoryCat)

		wordGivenCategory[cat] = wordGivenCategoryCat
	}
	log.Debug("category mapping now is:", category)
	log.Debug("Word given category mapping is:", wordGivenCategory)

	// so, no zero in denominator
	if totalCategories == 0 {
		totalCategories = 1
	}

	// evaluate result of training
	return func(words []string) (int, float64) {
		bestCatIndex := -1
		bestP := float64(math.MinInt32)
		log.Debug("Best P now is:", bestP)
		log.Debug("Calculating best category now")
		for index, catSlice := range input {

			totalP := 0.0
			catString := cat2String(catSlice)

			// P(Category)
			categoryAppearTimes := category[catString]
			log.Debug("Category appear times: ", categoryAppearTimes)
			pCat := math.Log(float64(categoryAppearTimes)) - math.Log(float64(totalCategories))
			log.Debug("pCat", pCat)
			totalP = pCat
			wordGivenCategory, ok := wordGivenCategory[catString]
			wordGivenCategoryLen := len(wordGivenCategory)
			if !ok {
				wordGivenCategory = map[string]uint32{}
			}

			log.Debug("Word Given Category ", wordGivenCategory)

			for _, wordItem := range words {
				// calculate P(Word|Category), P(Word, given Category)
				wordGivenCategoryAppearsTimes := wordGivenCategory[wordItem]
				pWordGivenCategory := math.Log(float64(wordGivenCategoryAppearsTimes)+laplase_nom) - math.Log(float64(wordGivenCategoryLen))
				log.Debug("P(Word, given category) =", pWordGivenCategory)
				totalP += pWordGivenCategory
			}
			log.Debug("Total P and best P and index are: ", totalP, bestP, index)

			if totalP > bestP {
				bestP = totalP
				bestCatIndex = index
			}
		}
		log.Debug("Best index and bestP are:", bestCatIndex, bestP)
		return bestCatIndex, bestP
	}
}
