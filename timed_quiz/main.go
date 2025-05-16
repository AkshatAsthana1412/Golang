package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file in the format 'qustion,answer'")
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
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer != p.a {
			fmt.Printf("Wrong answer %s, correct answer is %s\n", answer, p.a)
		} else {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
