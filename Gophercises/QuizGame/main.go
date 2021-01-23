package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type quizQuestion struct {
	question string
	answer   string
}

func main() {
	fileNameFlag := flag.String("cf", "problems.csv", "change quiz questions file, set to problems.csv by default.")
	timeLimitFlag := flag.String("t", "30s", "set the time limit for the quiz, set to 30s by default.")
	shuffleFlag := flag.Bool("shuffle", true, "decide whether the quiz questions should get shuffled with each run or not, set to true by default.")
	flag.Parse()

	if *shuffleFlag == true {
		rand.Seed(time.Now().UnixNano())
	}

	timeLimit, _ := time.ParseDuration(*timeLimitFlag)
	fmt.Println("time limit =", timeLimit)

	file, fileReadErr := os.Open(*fileNameFlag)
	if fileReadErr != nil {
		fmt.Println(fileReadErr)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true

	var quizQuestions []quizQuestion

	for {
		record, err := csvReader.Read()
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

	bufioReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything to start:")
	_, _, bufioReadErr := bufioReader.ReadRune()
	if bufioReadErr != nil {
		log.Fatal(bufioReadErr)
	}

	timeNow := time.Now()
	numCorrectAnswers := 0

	go func() {
		if *shuffleFlag == true {
			indices := rand.Perm(len(quizQuestions))
			for _, v := range indices {
				fmt.Println("time left:", timeLimit-time.Since(timeNow))
				askQuestion(quizQuestions[v], &numCorrectAnswers)
			}
		} else {
			for _, q := range quizQuestions {
				fmt.Println("time left:", timeLimit-time.Since(timeNow))
				askQuestion(q, &numCorrectAnswers)
			}
		}
		fmt.Println("You got", numCorrectAnswers, "correct questions out of", len(quizQuestions), "total questions.")
		os.Exit(0)
	}()

	for {
		if time.Since(timeNow) >= timeLimit {
			fmt.Println("Time is up!")
			fmt.Println("You got", numCorrectAnswers, "correct questions out of", len(quizQuestions), "total questions.")
			os.Exit(0)
		}
	}
}

func askQuestion(q quizQuestion, numCorrectAnswers *int) {
	var userAnswer string
	fmt.Println(q.question, "=")
	_, err := fmt.Scan(&userAnswer)
	if err != nil {
		log.Fatal("scanning answer error:", err)
	}
	if userAnswer == q.answer {
		*numCorrectAnswers++
	}
}
