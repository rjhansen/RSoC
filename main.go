package main

import (
	"fmt"
	"regexp"
	"strings"
)
import "io/ioutil"
import "log"

func main() {
	for _, v := range loadWords() {
		fmt.Println(v)

	}
}

func loadWords() []string {
	rawContent, err := ioutil.ReadFile("words")
	if err != nil {
		log.Fatal(err)
	}
	words := make([]string, 0)
	for _, v := range strings.Split(string(rawContent), "\n") {
		v = strings.TrimSpace(v)
		match, err := regexp.MatchString(`^[a-z]+$`, v)
		if err == nil && match {
			words = append(words, v)
		}
	}
	return words
}
