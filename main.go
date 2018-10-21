package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "problems.csv"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	questions, err := getQuestions(path)

	if err != nil {
		panic(err)
	}

	correctCount := 0

	for i, q := range questions {
		number := i + 1
		question := q[0]
		answer := q[1]

		printQuestion(number, question)

		a, err := getAnswer()

		if err != nil {
			panic(err)
		}

		correct := a == answer

		if correct {
			correctCount++
		}

		printStatus(correctCount, number)
	}

	fmt.Print("\n")
}

func getQuestions(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}

func printStatus(correct int, total int) {
	perc := float64(100 * correct) / float64(total)

	fmt.Printf("%d/%d (%.1f%%)\n", correct, total, perc)
}

func printQuestion(number int, question string) {
	fmt.Printf("Question #%d: %s = ", number, question)
}

func getAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	buffer, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer, "\n"), nil
}
