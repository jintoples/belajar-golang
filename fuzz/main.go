package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, err1 := Reverse2(input)
	doubleRev, err2 := Reverse2(rev)

	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, err1)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, err2)
}

func Reverse(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func Reverse2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}

	b := []rune(s)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}
