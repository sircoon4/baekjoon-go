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
		{input1: "3", input2: "3", input3: "6", output: "1300"},
		{input1: "2", input2: "2", input3: "2", output: "12000"},
		{input1: "6", input2: "2", input3: "5", output: "600"},
	}

	for i, test := range tests {
		expected := test.output
		cmd := exec.Command("go", "run", "main.go")
		cmd.Stdin = bytes.NewBufferString(test.input1 + " " + test.input2 + " " + test.input3)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			t.Fatalf("%v: Failed to run main.go: %v", i, err)
		}
		output := out.String()
		if output != expected {
			t.Errorf("%v: Expected: %v, Got: %v", i, expected, output)
		}
	}
}
