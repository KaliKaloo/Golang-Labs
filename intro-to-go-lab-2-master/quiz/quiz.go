package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

func ask(sChan chan score, question question) {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		sChan <- 1
	} else {
		fmt.Println("Incorrect :-(")
		sChan <- 0
	}
}

func main() {
	sChan := make(chan score)
	s := score(0)
	timer := time.After(5 * time.Second)
	qs := questions()

	for _, q := range qs {
		go ask(sChan, q) //ask and timer at same time as timer is also goroutine
		select {
		case r := <-sChan:
			{
				s += r
			}
		case <-timer:
			{
				fmt.Println("\nTimed out, final score", s)
				return
			}
		}
	}
}
