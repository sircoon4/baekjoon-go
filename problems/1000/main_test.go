package main

import (
	"bytes"
	"os/exec"
	"testing"
)

type testCase struct {
	input1 string
	input2 string
	output string
}

func TestMain(t *testing.T) {
	tests := []testCase{
		{input1: "1", input2: "2", output: "3"},
	}

	for _, test := range tests {
		expected := test.output
		cmd := exec.Command("go", "run", "main.go")
		cmd.Stdin = bytes.NewBufferString(test.input1 + " " + test.input2)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			t.Fatalf("Failed to run main.go: %v", err)
		}
		output := out.String()
		if output != expected {
			t.Errorf("Expected: %v, Got: %v", expected, output)
		}
	}
}
