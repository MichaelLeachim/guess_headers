// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 19:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	unidecode "github.com/mozillazg/go-unidecode"
	"regexp"
	"strings"
)

// Will transliterate unicode input into ascii input
func TokenizeUnidecode(input string) string {
	return unidecode.Unidecode(input)
}

func TokenizeLowercase(input string) string {
	return strings.ToLower(input)

}

// Will take only aphanumeric data from the input

var TokenizeAlphaNumericOnly = func() func(string) string {
	alphanum := regexp.MustCompile("[A-Za-z0-9]+")

	return func(input string) string {
		return strings.Join(alphanum.FindAllString(input, -1), " ")
	}
}()

// [Tokenize] Tokenize numbers via splitting, i.e. 1924 => (1000 900 20 4)
func TokenizeNumbers(number string) string {
	result := []string{}
	numLen := len(number)

	for index, number := range number {
		result = append(result, string(number)+strings.Repeat("0", numLen-index))
	}
	return strings.Join(result, " ")
}

func ShinglifyString(shingleSize int, data string) string {
	result := []string{}
	for i := 0; i <= len(data)-shingleSize; i++ {
		result = append(result, data[i:i+shingleSize])
	}
	return strings.Join(result, " ")
}

func ShinglifyTokenizedString(shingleSize int, data string) string {
	result := []string{}
	for _, item := range strings.Split(data, " ") {
		result = append(result, strings.TrimSpace(ShinglifyString(shingleSize, item)))
	}
	return strings.Join(result, " ")
}

func ShinglifyTokenizedString5(data string) string {
	return ShinglifyTokenizedString(5, data)
}
func ShinglifyTokenizedString3(data string) string {
	return ShinglifyTokenizedString(3, data)
}

func ShinglifyTokenizedString7(data string) string {
	return ShinglifyTokenizedString(7, data)
}

// will lose field information
func ReTokenizeOnSpace(data []string) []string {
	return strings.Split(strings.Join(data, " "), " ")
}

func ApplyRetokenizeOnSpaceToMatrix(data [][]string) [][]string {
	result := [][]string{}
	for _, row := range data {
		result = append(result, ReTokenizeOnSpace(row))
	}
	return result
}

func ApplyTokenizerToRow(data []string, workers ...func(string) string) []string {
	result := []string{}
	for _, item := range data {
		for _, worker := range workers {
			item = worker(item)
		}
		result = append(result, item)
	}
	return result
}

func ApplyTokenizerToMatrix(data [][]string, workers ...func(string) string) [][]string {
	result := [][]string{}
	for _, row := range data {
		result = append(result, ApplyTokenizerToRow(row, workers...))
	}
	return result
}
