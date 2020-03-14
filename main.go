package main

import (
	"fmt"
	"strings"
)
import "io/ioutil"
import "log"

func main() {
	rawContent, err := ioutil.ReadFile("words")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(rawContent), "\n")
	for _, v := range words {
		fmt.Println(v)
	}
}
