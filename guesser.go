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
	size = size - 1
	if size <= 0 {
		return []string{}
	}
	sizeofinput := int32(len(input))
	if sizeofinput == 0 {
		return []string{}
	}
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
	for index, input2 := range input2 {
		score := scoreFn(input1, input2)
		if best < score {
			best = score
			bestItem = Triplet{Left: input1, Right: input2, RightIndex: index, Score: best, Kind: TRIPLET_BOTH_MATCH}
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
		triplets = append(triplets, Triplet{Left: []string{}, Right: row, Score: 0, Kind: TRIPLET_RIGHT_ONLY})
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
			data[index].Kind = TRIPLET_LEFT_ONLY
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

// this is a base implementation of a guess function for row guessing
func BaseGuessRowsFunction(input, output [][]string, hasDuplicates bool) [][]string {
	// tokenize data
	headersOfInput, input := input[0], input[1:]
	headersOfOutput, output := output[0], output[1:]
	padInput := RepeatString("", len(headersOfInput))
	padOutput := RepeatString("", len(headersOfOutput))

	tokenizedInput := ApplyRetokenizeOnSpaceToMatrix(ApplyTokenizerToMatrix(input, TokenizeUnidecode, TokenizeLowercase, TokenizeNumbers, TokenizeAlphaNumericOnly))
	tokenizedOutput := ApplyRetokenizeOnSpaceToMatrix(ApplyTokenizerToMatrix(output, TokenizeUnidecode, TokenizeLowercase, TokenizeNumbers, TokenizeAlphaNumericOnly))

	// join two datasets (via Simple joiner) There are also: TfIdf and Naive Bayes joiners
	triplets := []Triplet{}
	for _, input := range tokenizedInput {
		triplets = append(triplets, CalculateBestMatch(MatchBetweenSimple, input, tokenizedOutput))
	}

	// remove duplicate matches
	if !hasDuplicates {
		triplets = CleanUp(triplets)
	}

	triplets = BuildUp(triplets, output)
	result := append([][]string{}, append(headersOfInput, headersOfOutput...))
	for index, triplet := range triplets {
		switch triplet.Kind {
		case TRIPLET_BOTH_MATCH:
			result = append(result, append(input[index], output[triplet.RightIndex]...))
		case TRIPLET_LEFT_ONLY:
			result = append(result, append(input[index], padOutput...))
		case TRIPLET_RIGHT_ONLY:
			result = append(result, append(padInput, output[triplet.RightIndex]...))
		}
	}
	return result
}

// this is a base implementation of a guess function for column guessing
func BaseGuessColumnsFunction(input, output [][]string) (map[string]string, []string, []string) {

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

	concordance := map[string]string{}
	onlyLeft := []string{}
	onlyRight := []string{}

	for _, triplet := range triplets {
		switch triplet.Kind {
		case TRIPLET_BOTH_MATCH:
			concordance[triplet.Left[0]] = triplet.Right[0]
		case TRIPLET_LEFT_ONLY:
			onlyLeft = append(onlyLeft, triplet.Left[0])
		case TRIPLET_RIGHT_ONLY:
			onlyRight = append(onlyRight, triplet.Right[0])
		}
	}
	return concordance, onlyLeft, onlyRight
}
