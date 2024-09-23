package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") // args:- name, default value, usage
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file."))
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("%d. %s\n", i+1, p.q)

		var answer string

		fmt.Scanf("%s\n", &answer)
		if answer != p.a {
			fmt.Println("Incorrect!")
			continue
		}

		fmt.Println("Correct!")
		correct++
	}

	fmt.Printf("Your score :- %d\n", correct)

}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines)) // creates a slice of type problem of the given length

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
