package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// An array of rules forms the rules of the game
type rule struct {
	Num int    `json:"num"`
	Wrd string `json:"wrd"`
}

// business logic : determines what string the number will map to based on rules
func fizzbuzz(n int, rules []rule) (s string) {
	// If empty rules, use default
	if len(rules) == 0 {
		log.Fatal("No rules")
	}

	for _, r := range rules {
		if n%r.Num == 0 {
			s += r.Wrd
		}
	}

	if s == "" {
		s = strconv.Itoa(n)
	}

	return
}

// given a path to a rules file, read it and return the rules as a []rule{}
// 	rules file (viz. rules.txt, default_rules.txt) will contain
//	a JSON formatted string of each rule in each line
// 	For Example :
// 		{"num": 3, "wrd": "Fizz"}
// 		{"num": 5, "wrd": "Buzz"}
func getRulesFromFile(pathToFile string) []rule {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rules []rule
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ruleStr := scanner.Text()
		var r rule
		err := json.Unmarshal([]byte(ruleStr), &r)
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rules
}
