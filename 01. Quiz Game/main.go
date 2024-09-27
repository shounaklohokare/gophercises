package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") // args:- name, default value, usage
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	problems := parseLines(shuffleLines(lines))

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // returns a channel when time is up

	correct := 0
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)

		go func() { // A closure that accepts the answer from user in a non-blocking manner
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {

		case <-timer.C: // checks if time is up
			fmt.Printf("You ran out of time. You scored %d out of %d\n", correct, len(problems))
			return

		case answer := <-answerCh: // checks if value is scanned in the closure

			if answer != p.a {
				fmt.Println("Incorrect!")
				continue
			}

			fmt.Println("Correct!")
			correct++
		}

	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines)) // creates a slice of type problem of the given length

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func shuffleLines(lines [][]string) [][]string { // shuffles lines to give a random order of questions each time

	r := rand.New(rand.NewSource(time.Now().Unix()))

	ret := make([][]string, len(lines))
	for i, randIndex := range r.Perm(len(lines)) {
		ret[i] = lines[randIndex]
	}

	return ret

}
