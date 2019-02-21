// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-19 19:55 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Let's simplify:
// We've got input, a csv file, like this

// [["name" "email"]
//  ["Michael" "thereisnodotcollective@gmail.com"]]

// Now, we can transpose this input, making it:
// [["name" "Michael"] ["email" "thereisnodotcollective@gmail.com"]]

// Now, we might take only a sample of this input, saving field names and then joining them again together
// And, we can also process the fields to make them easier to match via tokenizers.

// After that, we will have:
// * row to row match.
// * null on the right (not matched)
// * null on the left  (not matched)

// Then, we will be able to use those triplets to return data, either as:
// * headers of the first file, on the second file
// * reverse of
// * concordance between field names

// will return concordance between right and left headers of the dataset
func ReturnConcordHeaders(input, output [][]string) [][]string {
	return [][]string{}

}

// Will return right file with left headers
func ReturnRightFileWithLeftHeaders(data [][]string) [][]string {
	return [][]string{}

}

// Will return left file with right headers
func ReturnLeftFileWithRightHeaders(data [][]string) [][]string {
	return [][]string{}

}

// will return left join between two files
func ReturnLeftJoinBetweenFiles(data [][]string) [][]string {
	return [][]string{}

}
