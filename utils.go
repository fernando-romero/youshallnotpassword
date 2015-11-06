package main

import (
	"math/rand"
	"time"
)

func contains(slc []string, sub string) bool {
	for _, s := range slc {
		if s == sub {
			return true
		}
	}
	return false
}

func containsRune(runes []rune, r rune) bool {
	for _, ru := range runes {
		if ru == r {
			return true
		}
	}
	return false
}

func randomRune(runes []rune) rune {
	rand.Seed(time.Now().UTC().UnixNano())
	return runes[rand.Intn(len(runes))]
}

func isAlpha(r rune) bool {
	return containsRune(alphas, r)
}

func isDigit(r rune) bool {
	return containsRune(digits, r)
}

func isDelimiter(r rune) bool {
	return containsRune(delimiters, r)
}

func isOther(r rune) bool {
	return !isAlpha(r) && !isDigit(r) && !isDelimiter(r)
}
