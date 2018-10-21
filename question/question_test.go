package question

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestIsCorrect(t *testing.T) {
	in := Question{"question", "answer", 1, "answer"}
	out := in.IsCorrect()

	if !out {
		t.Errorf("Question{...}.IsCorrect() == %v, want %v", out, true)
	}
}

func TestNewQuestion(t *testing.T) {
	in := []string{"question", "answer"}
	out := NewQuestion(in, 1)

	if out.Prompt != "question" {
		t.Errorf("NewQuestion(...).Prompt == %v, want %v", out.Prompt, "question")
	}

	if out.ExpectedAnswer != "answer" {
		t.Errorf("NewQuestion(...).ExpectedAnswer == %v, want %v", out.ExpectedAnswer, "answer")
	}

	if out.Number != 1 {
		t.Errorf("NewQuestion(...).Number == %v, want %v", out.Number, 1)
	}

	if out.GivenAnswer != "" {
		t.Errorf("NewQuestion(...).GivenAnswer == %v, want %v", out.GivenAnswer, "")
	}
}

func TestPrintPrompt(t *testing.T) {
	in := Question{"question", "answer", 1, "answer"}
	buffer := bytes.Buffer{}
	writer := io.Writer(&buffer)
	reader := io.Reader(&buffer)

	in.PrintPrompt(writer)

	out, err := ioutil.ReadAll(reader)

	if err != nil {
		t.Error(err)
	}

	if string(out) != "Question #1: question = " {
		t.Errorf("Question{...}.PrintPrompt(...) == %v, want %v", string(out), "Question #1: question = ")
	}
}
