package main

import (
	"os"
	"testing"
)

func Test_exportContents(t *testing.T) {
	t.Setenv("example", "")
	exportContents("test")
	// Check if the environment variables are set
	if os.Getenv("example") == "ExampleContent" {
		t.Error("Environment variable test is not set")
	}
}
