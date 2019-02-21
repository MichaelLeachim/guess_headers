// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:09 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"math/rand"
	"strings"
)

// Guessing algorithm (ideas):
// [Reduce space] Take random N(100) fields of each column of each file (reducing step)
// [Tokenize] Tokenize numbers via splitting, i.e. 1924 => (1000 900 20 4)
// [Match] Match columns: cross join with each other, keep only the best match
// [Cleanup] Assume that there is only single one best match.
// If it is taken, there is nothing left (the match is nil)
// [BuildUp]   For those from the right which wasn't chosen by the left
// If it wasn't taken up by the match stage

// [Reduce amount] Take random N(100) fields of each column of each file (reducing step)
func TakeSeed(size int, input []string) []string {
	sizeofinput := int32(len(input))
	result := []string{}
	for i := 0; i <= size; i++ {
		result = append(result, input[rand.Int31n(sizeofinput)])
	}
	return result
}

// same as <TakeSeed>, but for the collection of columns
func TakeSeedOfList(size int, input [][]string) [][]string {
	for index, item := range input {
		input[index] = TakeSeed(size, item)
	}
	return input
}

func MatchBetweenSimple(input1, input2 []string) float64 {
	matches := 0
	for _, input1 := range input1 {
		for _, input2 := range input2 {
			if input1 == input2 {
				matches += 1
			}
		}
	}
	return float64(matches) / float64(((len(input1) + len(input2)) - matches))
}

// Chooses something from input2 that matches input1 in the best way
func CalculateBestMatch(scoreFn func([]string, []string) float64, input1 []string, input2 [][]string) Triplet {
	best := float64(0.0)
	bestItem := Triplet{}
	for _, input2 := range input2 {
		score := scoreFn(input1, input2)
		if best < score {
			best = score
			bestItem = Triplet{Left: input1, Right: input2, Score: best}
		}
	}
	return bestItem
}

// [BuildUp]   For those from the right which wasn't chosen by the left
// If it wasn't taken up by the match stage
func BuildUp(triplets []Triplet, rightSide [][]string) []Triplet {
	rightStore := map[string]bool{}

	for _, triplet := range triplets {
		rightStore[strings.Join(triplet.Right, "")] = true
	}
	for _, row := range rightSide {
		// already have this data matched
		if rightStore[strings.Join(row, "")] {
			continue
		}
		// this data is unmatched, so, add it
		triplets = append(triplets, Triplet{Left: nil, Right: row, Score: 0})
	}
	return triplets
}

// [Cleanup] Assume that there is only single best match. If it is taken, there is nothing left (the match is nil)
func CleanUp(data []Triplet) []Triplet {
	scores := map[string]float64{}
	for _, item := range data {
		score, ok := scores[strings.Join(item.Right, "")]
		if ok && (score <= item.Score) {
			scores[strings.Join(item.Right, "")] = item.Score
		}
	}
	for index, item := range data {
		score, ok := scores[strings.Join(item.Right, "")]
		if ok && (score <= item.Score) {
			data[index].Right = nil
			data[index].Score = 0
		}
	}
	return data
}

// will take headers of every column in a given input
func ChunkOffHeaders(input [][]string) ([]string, [][]string) {
	headers := []string{}
	body := [][]string{}

	for _, point := range input {
		headers = append(headers, point[0])
		body = append(body, point[0:])
	}
	return headers, body
}

// the opposite function of ChunkOffHeaders
func JoinUpHeaders(headers []string, body [][]string) [][]string {
	for index, point := range body {
		body[index] = append([]string{headers[index]}, point...)
	}
	return body
}

func Concordance(data []Triplet) [][]string {
	result := [][]string{}
	for _, item := range data {
		result = append(result, []string{item.Left[0], item.Right[0]})
	}
	return result
}

// this is a base implementation of a guess function for columns guessing
func BaseGuessColumnsFunction(input, output [][]string) []Triplet {

	// chunk off heads of every row
	inputHeaders, inputBody := ChunkOffHeaders(input)
	outputHeaders, outputBody := ChunkOffHeaders(output)

	// take seed of data
	inputSeed := TakeSeedOfList(100, inputBody)
	outputSeed := TakeSeedOfList(100, outputBody)

	// tokenize data
	tokenizedInputSeed := ApplyRetokenizeOnSpaceToMatrix(ApplyTokenizerToMatrix(inputSeed, TokenizeUnidecode, TokenizeLowercase, TokenizeNumbers, TokenizeAlphaNumericOnly))
	tokenizedOutputSeed := ApplyRetokenizeOnSpaceToMatrix(ApplyTokenizerToMatrix(outputSeed, TokenizeUnidecode, TokenizeLowercase, TokenizeNumbers, TokenizeAlphaNumericOnly))

	// append headers to tokenized data
	tokenizedInputSeedWithHeaders := JoinUpHeaders(inputHeaders, tokenizedInputSeed)
	tokenizedOutputSeedWithHeaders := JoinUpHeaders(outputHeaders, tokenizedOutputSeed)

	// join two datasets (via Simple joiner) There are also: TfIdf and Naive Bayes joiners
	triplets := []Triplet{}
	for _, input := range tokenizedInputSeedWithHeaders {
		triplets = append(triplets, CalculateBestMatch(MatchBetweenSimple, input, tokenizedOutputSeedWithHeaders))
	}

	// remove every right, if it is duplicated
	triplets = CleanUp(triplets)
	// append every left, that wasn't used
	triplets = BuildUp(triplets, tokenizedOutputSeedWithHeaders)
	return triplets
}
