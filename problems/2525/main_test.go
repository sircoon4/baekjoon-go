package main

import (
	"bytes"
	"os/exec"
	"testing"
)

type testCase struct {
	input1 string
	input2 string
	input3 string
	output string
}

func TestMain(t *testing.T) {
	tests := []testCase{
		{input1: "14", input2: "30", input3: "20", output: "14 50"},
		{input1: "17", input2: "40", input3: "80", output: "19 0"},
		{input1: "23", input2: "48", input3: "25", output: "0 13"},
	}

	for _, test := range tests {
		expected := test.output
		cmd := exec.Command("go", "run", "main.go")
		cmd.Stdin = bytes.NewBufferString(test.input1 + " " + test.input2 + "\n" + test.input3)
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
