package main

import (
	"flag"
	"fmt"
)

var rulesfile = flag.String("rules", "default_rules.txt", "Path to FizzBuzz game rules file")
var start = flag.Int("start", 1, "Value to start the game from")
var end = flag.Int("end", 100, "Value to End the game")

func main() {
	flag.Parse()

	// Get user defined rules; else default rules
	var rules []rule
	rules = getRulesFromFile(*rulesfile)

	for ii := *start; ii <= *end; ii++ {
		fmt.Println(fizzbuzz(ii, rules))
	}
}
