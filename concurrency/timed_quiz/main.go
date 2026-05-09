package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
	"timed_quiz/models"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file in the format 'qustion,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for each quiz in seconds")
	flag.Parse() // to parse all flags

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read %s", *csvFileName))
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer != p.A {
				fmt.Printf("Wrong answer %s, correct answer is %s\n", answer, p.A)
			} else {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []models.Problem {
	ret := make([]models.Problem, len(lines))
	for i, line := range lines {
		ret[i] = models.Problem{
			Q: line[0],
			A: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
