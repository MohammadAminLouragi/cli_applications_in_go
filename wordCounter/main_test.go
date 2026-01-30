package main

import (
	"bytes"
	"testing"
)

func TestCountWords(test *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6\n")

	exp := 6
	res := count(b)
	if res != exp {
		test.Fatalf("Expected %d, got %d instead. \n", exp, res)
	}
}
