// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:19 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTakeSeed(t *testing.T) {
	assert.Equal(t, TakeSeed(3, []string{}), []string{})
	assert.Equal(t, TakeSeed(3, []string{"A"}), []string{"A", "A", "A"})
	assert.Equal(t, len(TakeSeed(1000, []string{"A", "B", "N"})), 1000)

}

func TestBuildUp(t *testing.T) {
	triplets := []Triplet{
		newTestTriplet(1.0, "1912", "1912", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.5, "1912 blab", "1912 blob", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.0, "hello", "world", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.3, "whatever", "never", TRIPLET_BOTH_MATCH),
		newTestTriplet(1.0, "never", "never", TRIPLET_BOTH_MATCH),
		newTestTriplet(1.0, "alltimes", "alltimes", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.7, "sometimes", "alltimes", TRIPLET_BOTH_MATCH),
	}
	rightSide := [][]string{
		[]string{"1912"},
		[]string{"1912", "blob"},
		[]string{"world"},
		[]string{"never"},
		[]string{"alltimes"},
		[]string{"foo"},
		[]string{"bar"},
		[]string{"baz"},
	}
	resultTriplets := BuildUp(triplets, rightSide)
	assert.Equal(t, resultTriplets[7].Right, []string{"foo"})
	assert.Equal(t, resultTriplets[7].Kind, TRIPLET_RIGHT_ONLY)
	assert.Equal(t, resultTriplets[8].Right, []string{"bar"})
	assert.Equal(t, resultTriplets[8].Kind, TRIPLET_RIGHT_ONLY)
	assert.Equal(t, resultTriplets[9].Right, []string{"baz"})
	assert.Equal(t, resultTriplets[9].Kind, TRIPLET_RIGHT_ONLY)
}

func TestCleanUp(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	x := []Triplet{
		newTestTriplet(1.0, "1912", "1912", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.5, "1912 blab", "1912 blob", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.0, "hello", "world", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.3, "whatever", "never", TRIPLET_BOTH_MATCH),
		newTestTriplet(1.0, "never", "never", TRIPLET_BOTH_MATCH),
		newTestTriplet(1.0, "alltimes", "alltimes", TRIPLET_BOTH_MATCH),
		newTestTriplet(0.7, "sometimes", "alltimes", TRIPLET_BOTH_MATCH),
	}
	clap := CleanUp(x)
	assert.Equal(t, clap[0].Kind, TRIPLET_BOTH_MATCH)
	assert.Equal(t, clap[1].Kind, TRIPLET_BOTH_MATCH)
	assert.Equal(t, clap[2].Kind, TRIPLET_BOTH_MATCH)
	assert.Equal(t, clap[3].Kind, TRIPLET_LEFT_ONLY)
	assert.Equal(t, clap[4].Kind, TRIPLET_BOTH_MATCH)
	assert.Equal(t, clap[5].Kind, TRIPLET_BOTH_MATCH)
	assert.Equal(t, clap[6].Kind, TRIPLET_LEFT_ONLY)
}

func TestMatchBetweenSimple(t *testing.T) {
	assert.Equal(t, MatchBetweenSimple([]string{"hello", "world", "foo", "bar"}, []string{"hello", "world"}), 0.5)
	assert.Equal(t, MatchBetweenSimple([]string{"A", "B", "C"}, []string{"A", "B"}), 0.6666666666666666)
	assert.Equal(t, MatchBetweenSimple([]string{}, []string{"hello", "world"}), 0.0)
	assert.Equal(t, MatchBetweenSimple([]string{}, []string{}), 1.0)
	assert.Equal(t, MatchBetweenSimple([]string{"A"}, []string{"B"}), 0.0)
}

func TestCalculateBestMatch(t *testing.T) {
	assert.Equal(t, CalculateBestMatch(MatchBetweenSimple, []string{"Hello", "World"}, [][]string{
		[]string{"Hello"},
		[]string{"World"},
		[]string{"Hello", "World", "Blab"},
		[]string{"Blab"},
		[]string{"Hello", "World"},
		[]string{""},
		[]string{},
	}).RightIndex, 4)
}

func TestChunkOffHeaders(t *testing.T) {
	headers, body := ChunkOffHeaders([][]string{})
	assert.Equal(t, headers, []string{})
	assert.Equal(t, body, [][]string{})

	// headers, body = ChunkOffHeaders([][]string{[]string{"Ha"}, []string{"Ho"}})
	// assert.Equal(t, headers, []string{"Ha", "Ho"})
	// assert.Equal(t, body, [][]string{})

	headers, body = ChunkOffHeaders([][]string{
		[]string{"name", "Michael"},
		[]string{"surname", "Leahcim"},
		[]string{"NoBody"},
		[]string{}})

	assert.Equal(t, headers, []string{"name", "surname", "NoBody", ""})
	assert.Equal(t, body, [][]string{[]string{"Michael"}, []string{"Leahcim"}, []string{}, []string{}})
}

func TestJoinUpHeaders(t *testing.T) {
	data := [][]string{
		[]string{"name", "Michael"},
		[]string{"surname", "Leahcim"},
		[]string{"NoBody"},
		[]string{}}
	headers, body := ChunkOffHeaders(data)
	assert.Equal(t, JoinUpHeaders(headers, body), data)

	data = [][]string{}
	headers, body = ChunkOffHeaders(data)
	assert.Equal(t, JoinUpHeaders(headers, body), data)
}

func TestBaseGuessRowsOnSmallRealWorldData(t *testing.T) {
	log.SetLevel(log.FatalLevel)
	countriesInRU, err := ReadCSVFile("testdata/countries.ru.csv", ',')
	assert.Equal(t, err, nil)
	countriesInEN, err := ReadCSVFile("testdata/countries.en.csv", ',')
	assert.Equal(t, err, nil)
	countriesInRU = append(countriesInRU, []string{"Country Name"})
	countriesInEN = append(countriesInEN, []string{"Country Name"})
	assert.Equal(t, len(countriesInRU), 211)
	assert.Equal(t, len(countriesInEN), 211)
	matches := func(bgrf [][]string) int {
		matches := 0
		for index, item := range bgrf {
			guessedRU, guessedEN := item[0], item[1]
			if guessedEN != "" && guessedRU != "" && countriesInEN[index][0] == guessedEN && countriesInRU[index][0] == guessedRU {
				matches += 1
			}
		}
		return matches
	}
	assert.Equal(t, matches(BaseGuessRowsFunction(countriesInRU, countriesInEN, MATCH_BAYES, false)), 153)
	assert.Equal(t, matches(BaseGuessRowsFunction(countriesInRU, countriesInEN, MATCH_SIMPLE, false)), 150)
	assert.Equal(t, matches(BaseGuessRowsFunction(countriesInRU, countriesInEN, MATCH_TFIDF, false)), 62)

}

func TestBaseGuessRowsFunction(t *testing.T) {
	assert.Equal(t,
		[][]string{
			[]string{"City or Country Name in EN", "City or Country Name in RU"},
			[]string{"Petersburg", "Peterburgh"},
			[]string{"Moscow", "Moskva"},
			[]string{"France", "Frantsia"},
			[]string{"Lebanon", ""},
			[]string{"Bulgaria", "Bolgaria"},
			[]string{"", "Livan"}},
		BaseGuessRowsFunction([][]string{
			[]string{"City or Country Name in EN"},
			[]string{"Petersburg"},
			[]string{"Moscow"},
			[]string{"France"},
			[]string{"Lebanon"},
			[]string{"Bulgaria"},
		}, [][]string{
			[]string{"City or Country Name in RU"},
			[]string{"Moskva"},
			[]string{"Peterburgh"},
			[]string{"Frantsia"},
			[]string{"Livan"},
			[]string{"Bolgaria"},
		}, MATCH_SIMPLE, false))

	assert.Equal(t, BaseGuessRowsFunction([][]string{}, [][]string{}, MATCH_SIMPLE, false), [][]string{})

	assert.Equal(t,
		[][]string{
			[]string{"City Name", "City Name"},
			[]string{"Moscow", "Moscow"},
			[]string{"Moskva", "Moscow"},
			[]string{"Maskov", "Moscow"},
			[]string{"Sofia", "Moscow"},
			[]string{"Sofia", "Moscow"},
			[]string{"Sofia", "Moscow"},
			[]string{"", "Warshaw"},
			[]string{"", "New-York"},
			[]string{"", "Berlin"}},
		BaseGuessRowsFunction([][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moskva"},
			[]string{"Maskov"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, [][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Warshaw"},
			[]string{"New-York"},
			[]string{"Berlin"},
		}, MATCH_SIMPLE, true))
	assert.Equal(t,
		[][]string{[]string{"City Name", "City Name"},
			[]string{"Moscow", "Moscow"},
			[]string{"Warshaw", ""},
			[]string{"New-York", ""},
			[]string{"Berlin", ""},
			[]string{"", "Sofia"}},

		BaseGuessRowsFunction([][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Warshaw"},
			[]string{"New-York"},
			[]string{"Berlin"},
		}, [][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, MATCH_SIMPLE, false))

	assert.Equal(t, BaseGuessRowsFunction([][]string{},
		[][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, MATCH_SIMPLE, false), [][]string{})
	assert.Equal(t, BaseGuessRowsFunction(
		[][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, [][]string{}, MATCH_SIMPLE, false), [][]string{})
}

func TestBaseGuessColumnsFunction(t *testing.T) {
	csv1, err := ReadCSVFile("testdata/input1.csv", ',')
	csv2, err := ReadCSVFile("testdata/input2.csv", ',')
	assert.Equal(t, err, nil)
	assert.Equal(t, err, nil)
	concordace, left, right := BaseGuessColumnsFunction(TransposeMatrix(csv1), TransposeMatrix(csv2))

	assert.Equal(t, concordace, "")
	assert.Equal(t, left, "")
	assert.Equal(t, right, "")

}
