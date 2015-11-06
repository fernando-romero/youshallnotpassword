package main

import (
	"sort"
	"strings"
)

type Password string

func (password Password) Substrings() []string {
	var substrings []string
	passwordStr := string(password)
	for i := 0; i <= len(passwordStr); i++ {
		for j := i; j <= len(passwordStr); j++ {
			s := passwordStr[i:j]
			s = strings.ToLower(s)
			if !contains(substrings, s) {
				substrings = append(substrings, s)
			}
		}
	}
	sort.Sort(SliceOfStrings(substrings))
	return substrings
}

func (password Password) Words() []string {
	substrings := password.Substrings()
	var words []string
	for _, s := range substrings {
		if contains(dictionary, s) {
			words = append(words, s)
		}
	}
	return words
}

func (password Password) ReplaceWords() Password {
	passwordStr := string(password)
	words := password.Words()
	for _, w := range words {
		if strings.Contains(passwordStr, w) {
			r := randomRune(lowerLetters)
			s := string(r)
			passwordStr = strings.Replace(passwordStr, w, s, 1)
		}
	}
	return Password(passwordStr)
}

func (password Password) Strength() int {
	password = password.ReplaceWords()
	charTypes := 0
	if password.AlphaCount() > 0 {
		charTypes++
	}
	if password.DigitCount() > 0 {
		charTypes++
	}
	if password.DelimiterCount() > 0 {
		charTypes++
	}
	if password.OtherCount() > 0 {
		charTypes++
	}
	return len(string(password)) * charTypes
}

type IsTypeFn func(rune) bool

func (password Password) ReplaceFirst(fn IsTypeFn, r rune) Password {
	passwordStr := string(password)
	runes := []rune(passwordStr)
	for _, ru := range runes {
		if fn(ru) {
			passwordStr = strings.Replace(passwordStr, string(ru), string(r), 1)
			return Password(passwordStr)
		}
	}
	return password
}

func (password Password) Strengthen() Password {

	alphaCount := password.AlphaCount()
	digitCount := password.DigitCount()
	delimiterCount := password.DelimiterCount()
	otherCount := password.OtherCount()

	if alphaCount == 0 {
		r := randomRune(alphas)
		if digitCount > 1 {
			return password.ReplaceFirst(isDigit, r)
		}
		if delimiterCount > 1 {
			return password.ReplaceFirst(isDelimiter, r)
		}
		if otherCount > 1 {
			return password.ReplaceFirst(isOther, r)
		}
	}

	if digitCount == 0 {
		r := randomRune(digits)
		if alphaCount > 1 {
			return password.ReplaceFirst(isAlpha, r)
		}
		if delimiterCount > 1 {
			return password.ReplaceFirst(isDelimiter, r)
		}
		if otherCount > 1 {
			return password.ReplaceFirst(isOther, r)
		}
	}

	if delimiterCount == 0 {
		r := randomRune(delimiters)
		if alphaCount > 1 {
			return password.ReplaceFirst(isAlpha, r)
		}
		if digitCount > 1 {
			return password.ReplaceFirst(isDigit, r)
		}
		if otherCount > 1 {
			return password.ReplaceFirst(isOther, r)
		}
	}

	if otherCount == 0 {
		r := randomRune(others)
		if alphaCount > 1 {
			return password.ReplaceFirst(isAlpha, r)
		}
		if digitCount > 1 {
			return password.ReplaceFirst(isDigit, r)
		}
		if delimiterCount > 1 {
			return password.ReplaceFirst(isDelimiter, r)
		}
	}

	return password.AddChar()
}

func (password Password) AddChar() Password {
	runeToAdd := string(randomRune(characters))
	passwordStr := string(password)
	passwordStr = passwordStr + string(runeToAdd)
	return Password(passwordStr)
}

func (password Password) AlphaCount() int {
	runes := []rune(string(password))
	c := 0
	for _, r := range runes {
		if isAlpha(r) {
			c++
		}
	}
	return c
}

func (password Password) DigitCount() int {
	runes := []rune(string(password))
	c := 0
	for _, r := range runes {
		if isDigit(r) {
			c++
		}
	}
	return c
}

func (password Password) DelimiterCount() int {
	runes := []rune(string(password))
	c := 0
	for _, r := range runes {
		if isDelimiter(r) {
			c++
		}
	}
	return c
}

func (password Password) OtherCount() int {
	runes := []rune(string(password))
	c := 0
	for _, r := range runes {
		if isOther(r) {
			c++
		}
	}
	return c
}
