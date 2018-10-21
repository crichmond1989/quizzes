package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/crichmond1989/quizzes/quiz"
)

func main() {
	path := "problems.csv"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	quiz, err := NewQuizFromCSV(path)

	if err != nil {
		panic(err)
	}

	correctCount := 0

	for _, q := range quiz.Questions {
		q.PrintPrompt(os.Stdout)

		a, err := getAnswer()

		if err != nil {
			panic(err)
		}

		correct := a == q.ExpectedAnswer

		var resultIcon string

		if correct {
			correctCount++
			resultIcon = "\u2713"
		} else {
			resultIcon = "\u2717"
		}

		status := getStatus(correctCount, q.Number)

		fmt.Printf("%s %s", resultIcon, status)
	}

	fmt.Print("\n")
}

func getStatus(correct int, total int) string {
	perc := float64(100*correct) / float64(total)

	return fmt.Sprintf("%d/%d (%.1f%%)\n", correct, total, perc)
}

func getAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	buffer, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer, "\n"), nil
}
