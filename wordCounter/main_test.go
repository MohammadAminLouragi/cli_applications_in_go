package main

import (
	"bytes"
	"testing"
)

func TestCountWords(test *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6\n")

	exp := 6
	res := count(b, false)
	if res != exp {
		test.Fatalf("Expected %d, got %d instead. \n", exp, res)
	}
}

func TestCountLines(test *testing.T) {
	b := bytes.NewBufferString("word1 word2\n word3 \n word4 word5 word6")
	exp := 3
	res := count(b, true)
	if res != exp {
		test.Fatalf("Expected %d, got %d instead. \n", exp, res)
	}
}
