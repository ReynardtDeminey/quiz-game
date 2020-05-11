package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "The CSV file containing the questions and answers in the format 'question, answer'")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *fileName))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Could not parse the provided CSV file"))
	}
	problems := parseLines(lines)
	fmt.Println(problems)

	score := 0
	for i, p := range problems {
		fmt.Printf("Problem %d: %s= \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d correct!\n", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
