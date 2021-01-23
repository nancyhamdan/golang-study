package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type quizQuestion struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("cf", "problems.csv", "change quiz questions file")
	timeLimitFlag := flag.String("t", "30s", "set the time limit for the quiz")
	flag.Parse()

	timeLimit, _ := time.ParseDuration(*timeLimitFlag)
	fmt.Println("time limit = ", timeLimit)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true

	var quizQuestions []quizQuestion

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		question := record[0]
		answer := record[1]

		q := quizQuestion{question, answer}
		quizQuestions = append(quizQuestions, q)
	}

	noCorrectAnswers := 0

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything to start:")
	_, _, rerr := reader.ReadRune()
	if rerr != nil {
		log.Fatal(rerr)
	}

	timeNow := time.Now()

	go func() {
		for i, v := range quizQuestions {
			fmt.Println("time since = ", time.Since(timeNow))
			fmt.Println("Problem #", i+1, ":", v.question, "=")

			var userAnswer string

			_, err := fmt.Scan(&userAnswer)
			if err != nil {
				log.Fatal("scanning answer error:", err)
			}
			if userAnswer == v.answer {
				noCorrectAnswers++
			}
		}
	}()

	for {
		if time.Since(timeNow) >= timeLimit {
			fmt.Println("Time is up!")
			fmt.Println(noCorrectAnswers, "correct questions out of", len(quizQuestions), "total questions.")
			os.Exit(0)
		}
	}
}
