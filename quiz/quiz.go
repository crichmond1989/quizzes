package quiz

import (
	"encoding/csv"
	. "github.com/crichmond1989/quizzes/question"
	"os"
)

type Quiz struct {
	Questions []Question
}

func NewQuiz(source [][]string) Quiz {
	items := []Question{}

	for i, q := range source {
		item := NewQuestion(q, i+1)
		items = append(items, item)
	}

	return Quiz{items}
}

func NewQuizFromCSV(source string) (Quiz, error) {
	file, err := os.Open(source)

	if err != nil {
		return Quiz{}, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return Quiz{}, err
	}

	return NewQuiz(lines), nil
}
