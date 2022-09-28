package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)


func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	timer := flag.Int("timer", 30, "time that you have to answer the questions")
	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	csv, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}


	problems := make([]problem, len(csv))
	t := time.Now()
	correct := 0
	for i, p := range csv {
		var response string

		
		fmt.Printf("what is the answer to %s? ", p[0])

		fmt.Scanln(&response)

		problems[i] = problem{
			question: p[0],
			answer: p[1],
		}

		if response == p[1] {
			correct++
		}

		// TODO: finish program when timer hits `timer`
		now := time.Now().UTC().Second() - t.UTC().Second() 
		if now > *timer {
			break
		}
	}

	fmt.Printf("you got correct %d of %d questions\n", correct, len(csv))
}

type problem struct {
	question string
	answer string
}

