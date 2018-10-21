package quiz

import "testing"

func TestNewQuiz(t *testing.T) {
	in := [][]string{{"question", "answer"}}
	out := NewQuiz(in)

	if len(out.Questions) != 1 {
		t.Errorf("len(NewQuiz(...).Questions) == %v, want %v", len(out.Questions), 1)
	}
}
