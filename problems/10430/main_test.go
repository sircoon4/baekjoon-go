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
		{input1: "5", input2: "8", input3: "4", output: "1\n1\n0\n0"},
	}

	for _, test := range tests {
		expected := test.output
		cmd := exec.Command("go", "run", "main.go")
		cmd.Stdin = bytes.NewBufferString(test.input1 + " " + test.input2 + " " + test.input3)
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
