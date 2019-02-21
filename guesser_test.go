// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:19 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTakeSeed(t *testing.T) {
	assert.Equal(t, TakeSeed(3, []string{}), "")
	assert.Equal(t, TakeSeed(3, []string{"A"}), []string{"A", "A", "A"})
	assert.Equal(t, len(TakeSeed(1000, []string{"A", "B", "N"})), 1000)
}

func TestBuildUp(t *testing.T) {
	triplets := []Triplet{
		newTestTriplet(1.0, "1912", "1912"),
		newTestTriplet(0.5, "1912 blab", "1912 blob"),
		newTestTriplet(0.0, "hello", "world"),
		newTestTriplet(0.3, "whatever", "never"),
		newTestTriplet(1.0, "never", "never"),
		newTestTriplet(1.0, "alltimes", "alltimes"),
		newTestTriplet(0.7, "sometimes", "alltimes"),
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
	resultTriplets[8].Left = []string{"bar"}
	resultTriplets[9].Left = []string{"baz"}

}

func TestCleanUp(t *testing.T) {
	x := []Triplet{
		newTestTriplet(1.0, "1912", "1912"),
		newTestTriplet(0.5, "1912 blab", "1912 blob"),
		newTestTriplet(0.0, "hello", "world"),
		newTestTriplet(0.3, "whatever", "never"),
		newTestTriplet(1.0, "never", "never"),
		newTestTriplet(1.0, "alltimes", "alltimes"),
		newTestTriplet(0.7, "sometimes", "alltimes"),
	}
	assert.Equal(t, CleanUp(x), "")
}

func TestMatchBetweenSimple(t *testing.T) {
	assert.Equal(t, MatchBetweenSimple([]string{"hello", "world", "foo", "bar"}, []string{"hello", "world"}), "")
	assert.Equal(t, MatchBetweenSimple([]string{}, []string{"hello", "world"}), "")
	assert.Equal(t, MatchBetweenSimple([]string{}, []string{}), "")
	assert.Equal(t, MatchBetweenSimple([]string{"A"}, []string{"B"}), "")
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
