package main

import (
	"bytes"
	"testing"
)

func TestCountWords(test *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6\n")

	exp := 6
	res := count(b, false, false)
	if res != exp {
		test.Fatalf("Expected %d words, got %d instead. \n", exp, res)
	}
}

func TestCountLines(test *testing.T) {
	b := bytes.NewBufferString("word1 word2\n word3 \n word4 word5 word6")
	exp := 3
	res := count(b, true, false)
	if res != exp {
		test.Fatalf("Expected %d lines, got %d instead. \n", exp, res)
	}
}

func TestCountBytes(test *testing.T) {
	b := bytes.NewBufferString("word1 word2")
	exp := 11
	res := count(b, false, true)
	if res != exp {
		test.Fatalf("Expected %d bytes, got %d instead. \n", exp, res)
	}
}
