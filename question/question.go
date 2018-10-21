package question

import (
	"fmt"
	"io"
)

type Question struct {
	Prompt         string
	ExpectedAnswer string
	Number         int
	GivenAnswer    string
}

func NewQuestion(source []string, number int) Question {
	return Question{source[0], source[1], number, ""}
}

func (q *Question) IsCorrect() bool {
	return q.GivenAnswer == q.ExpectedAnswer
}

func (q *Question) PrintPrompt(w io.Writer) {
	fmt.Fprintf(w, "Question #%d: %s = ", q.Number, q.Prompt)
}
