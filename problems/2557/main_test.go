package main

import (
	"bytes"
	"os/exec"
	"testing"
)

type testCase struct {
	output string
}

func TestMain(t *testing.T) {
	tests := []testCase{
		{output: "Hello World!"},
	}

	for _, test := range tests {
		expected := test.output
		cmd := exec.Command("go", "run", "main.go")
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
