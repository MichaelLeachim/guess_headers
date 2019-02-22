// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 17:38 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMakeBayessianMatcher(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	splitstr := func(in string) []string {
		return strings.Split(in, " ")
	}
	data := [][]string{
		splitstr("Dogs are awesome, cats too. I love my dog"),
		splitstr("Cats are more preferred by software developers. I never could stand cats. I have a dog"),
		splitstr("My dog's name is Willy. He likes to play with my wife's cat all day long. I love dogs"),
		splitstr("Cats are difficult animals, unlike dogs, really annoying, I hate them all"),
		splitstr("So which one should you choose? A dog, definitely."),
		splitstr("The favorite food for cats is bird meat, although mice are good, but birds are a delicacy"),
		splitstr("A dog will eat anything, including birds or whatever meat"),
		splitstr("My cat's favorite place to purr is on my keyboard"),
		splitstr("My dog's favorite place to take a leak is the tree in front of our house"),
	}
	bayessianMatcher := makeBayessianMatcher(data)

	// basic matcher. Should return results if three words randomly taken from the base dataset
	for _, item := range data {
		matches, _ := bayessianMatcher(append(TakeSeed(3, item)))
		assert.NotEqual(t, matches, -1)
		assert.Equal(t, item, data[matches])
	}

	// should match, given 5 words from given category, and 3 words from another category
	for _, item := range data {
		matches, _ := bayessianMatcher(append(TakeSeed(5, item), TakeSeed(3, data[1])...))
		assert.NotEqual(t, matches, -1)
		assert.Equal(t, item, data[matches])
	}

	bayessianMatcher = makeBayessianMatcher([][]string{})
	for _, item := range data {
		index, _ := bayessianMatcher(item)
		assert.Equal(t, index, -1)
	}

}
