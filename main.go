package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func loadDictionary() ([]string, error) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	return words, nil
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s ByLength) Contains(sub string) bool {
	for _, a := range s {
		if a == sub {
			return true
		}
	}
	return false
}

func contains(slc []string, sub string) bool {
	for _, s := range slc {
		if s == sub {
			return true
		}
	}
	return false
}

func filterBlanksAndDups(slc []string) []string {
	var filtered []string
	for _, s := range slc {
		if s != "" && !contains(filtered, s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func getSubstrings(s string) []string {
	log.Println("Substrings")
	var substrings []string
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			substring := s[i:j]
			substring = strings.ToLower(substring)
			substrings = append(substrings, substring)
		}
	}
	substrings = filterBlanksAndDups(substrings)
	sort.Sort(ByLength(substrings))
	return substrings
}

func getWords(dictionary []string, slc []string) []string {
	var words []string
	for _, s := range slc {
		if contains(dictionary, s) {
			if len(s) > 2 {
				words = append(words, s)
			}
		}
	}
	return words
}

func getRandomLetter() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UTC().UnixNano())
	r := letterRunes[rand.Intn(len(letterRunes))]
	return string(r)
}

func calculateStrength(password string) int {
	runes := []rune(password)
	hasAlpha := false
	hasDigit := false
	hasDelimiter := false
	hasOther := false
	for _, r := range runes {
		if r >= 65 && r <= 90 {
			hasAlpha = true
		} else if r >= 48 && r <= 57 {
			hasDigit = true
		} else if r == 32 || r == 9 || r == 10 {
			hasDelimiter = true
		} else {
			hasOther = true
		}
	}
	strength := 0
	if hasAlpha {
		strength = strength + len(password)
	}
	if hasDigit {
		strength = strength + len(password)
	}
	if hasDelimiter {
		strength = strength + len(password)
	}
	if hasOther {
		strength = strength + len(password)
	}
	return strength
}

func main() {
	password := os.Args[1]
	dictionary, err := loadDictionary()
	if err != nil {
		log.Fatal(err)
	}
	substrings := getSubstrings(password)
	words := getWords(dictionary, substrings)
	for _, w := range words {
		log.Println("word:   " + w)
		if strings.Contains(password, w) {
			letter := getRandomLetter()
			log.Println("letter: " + letter)
			password = strings.Replace(password, w, letter, -1)
		}
	}
	log.Println("New password:")
	log.Println(password)
	strength := calculateStrength(password)
	log.Printf("Strength: %d", strength)
}
