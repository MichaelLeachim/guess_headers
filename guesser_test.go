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

func TestTakeSeed(t *testing.T){
	assert.Equal(t,TakeSeed(3, []string{}),"")
	assert.Equal(t,TakeSeed(3, []string{"A"}),[]string{"A","A","A"})
	assert.Equal(t,len(TakeSeed(1000, []string{"A","B","N"})),1000)
}

func TestCleanUp(t *testing.T)  {
	x := []Triplet{
		newTestTriplet(1.0,"1912","1912"),
		newTestTriplet(0.5,"1912 blab","1912 blob"),
		newTestTriplet(0.0,"hello","world"),
		newTestTriplet(0.3,"whatever","never"),
		newTestTriplet(1.0,"never","never"),
		newTestTriplet(1.0,"alltimes","alltimes"),
		newTestTriplet(0.7,"sometimes","alltimes"),
	}
	assert.Equal(t,CleanUp(x),"")
}

func TestMatchBetweenSimple(t *testing.T){
	assert.Equal(t,MatchBetweenSimple([]string{"hello","world","foo","bar"}, []string{"hello","world"}),"")
	assert.Equal(t,MatchBetweenSimple([]string{}, []string{"hello","world"}),"")
	assert.Equal(t,MatchBetweenSimple([]string{}, []string{}),"")
	assert.Equal(t,MatchBetweenSimple([]string{"A"}, []string{"B"}),"")
}

func TestCalculateBestMatch(t *testing.T){
	assert.Equal(t,CalculateBestMatch(MatchBetweenSimple, []string{"Hello", "World"}, [][]string{
		[]string{"Hello"},
		[]string{"World"},
		[]string{"Hello","World","Blab"},
		[]string{"Blab"},
		[]string{"Hello","World"},
		[]string{""},
		[]string{},
	}),"")
}

func TestSimpleGuessing(t *testing.T) {

	file1 := [][]string{
		[]string{}
		
	}

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}
