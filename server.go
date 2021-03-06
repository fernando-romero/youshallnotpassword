package main

import (
	"bufio"
	"log"
	"os"
)

var (
	characters   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 _-.!@#$%&?¿*")
	alphas       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
	digits       = []rune("0123456789")
	others       = []rune(".!@#$%&?¿*")
	delimiters   = []rune(" _-")
	dictionary   []string
)

func loadDictionary() {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		word := scanner.Text()
		dictionary = append(dictionary, word)
	}
}

func init() {
	loadDictionary()
}

func main() {
	password := Password(os.Args[1])
	strength := password.Strength()
	log.Printf("PASSWORD: %s", string(password))
	log.Printf("STRENGTH: %d", strength)
	if strength >= 50 {
		log.Println("CATEGORY: STRONG")
		log.Println("MESSAGE:  Congratulations, you shall password!")
		return
	}
	if strength <= 10 {
		log.Println("CATEGORY: UNACCEPTABLE")
		log.Println("MESSAGE:  You shall not password! Please try again with a better password")
		return
	}
	log.Println("CATEGORY: WEAK")
	log.Println("MESSAGE:  There is still hope, we will give you a better password here")
	for strength < 50 {
		password = password.Strengthen()
		strength = password.Strength()
	}
	log.Printf("PASSWORD: %s", string(password))
	log.Printf("STRENGTH: %d", strength)
}
