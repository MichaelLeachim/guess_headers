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
	resultTriplets[7].Left = []string{"foo"}
	resultTriplets[7].Kind = TRIPLET_RIGHT_ONLY
	resultTriplets[8].Left = []string{"bar"}
	resultTriplets[8].Kind = TRIPLET_RIGHT_ONLY
	resultTriplets[9].Left = []string{"baz"}
	resultTriplets[9].Kind = TRIPLET_RIGHT_ONLY
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
	}), "")
}

func TestChunkOffHeaders(t *testing.T) {
	headers, body := ChunkOffHeaders([][]string{})
	assert.Equal(t, headers, []string{})
	assert.Equal(t, body, [][]string{})

	headers, body = ChunkOffHeaders([][]string{
		[]string{"name", "Michael"},
		[]string{"surname", "Leahcim"},
		[]string{"NoBody"},
		[]string{}})

	assert.Equal(t, headers, []string{})
	assert.Equal(t, body, [][]string{})
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

func TestBaseGuessRowsFunction(t *testing.T) {
	assert.Equal(t, BaseGuessRowsFunction([][]string{
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
	}, false), "")
	countriesInRU, err := ReadCSVFile("testdata/countries.ru.csv", ',')
	assert.Equal(t, err, nil)
	countriesInEN, err := ReadCSVFile("testdata/countries.en.csv", ',')
	assert.Equal(t, err, nil)
	assert.Equal(t, BaseGuessRowsFunction(countriesInRU, countriesInEN, false), "")
	assert.Equal(t, BaseGuessRowsFunction([][]string{}, [][]string{}, false), "")
	assert.Equal(t, BaseGuessRowsFunction([][]string{
		[]string{"City Name"},
		[]string{"Moscow"},
		[]string{"Moscow"},
		[]string{"Moscow"},
		[]string{"Sofia"},
		[]string{"Sofia"},
		[]string{"Sofia"},
	}, [][]string{
		[]string{"City Name"},
		[]string{"Moscow"},
		[]string{"Warshaw"},
		[]string{"New-York"},
		[]string{"Berlin"},
	}, false), "")
	assert.Equal(t, BaseGuessRowsFunction([][]string{
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
	}, false), "")

	assert.Equal(t, BaseGuessRowsFunction([][]string{},
		[][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, false), "")
	assert.Equal(t, BaseGuessRowsFunction(
		[][]string{
			[]string{"City Name"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Moscow"},
			[]string{"Sofia"},
			[]string{"Sofia"},
			[]string{"Sofia"},
		}, [][]string{}, false), "")

	// TODO: test for empty
	// TODO: test for duplicates (on right) and on left
	// TODO: test for non comparables (zero score)
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
