// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:33 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

type Triplet struct {
	Left       []string
	Right      []string
	RightIndex int
	Score      float64
	Kind       int
}

const (
	TRIPLET_LEFT_ONLY = 1 + iota
	TRIPLET_RIGHT_ONLY
	TRIPLET_BOTH_MATCH
)

const (
	MATCH_BAYES = 1 + iota
	MATCH_SIMPLE
	MATCH_TFIDF
)
