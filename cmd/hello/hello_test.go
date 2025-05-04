package main

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/fatjyc/hello-go/pkg/hello"
)

func TestMain(t *testing.T) {
	// Save original flags
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test cases
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "default values",
			args:     []string{"cmd"},
			expected: "Hello, World![dlroW]\n",
		},
		{
			name:     "custom name",
			args:     []string{"cmd", "-name=John"},
			expected: "Hello, John![nhoJ]\n",
		},
		{
			name:     "custom language",
			args:     []string{"cmd", "-lang=es"},
			expected: "Hola, World![dlroW]\n",
		},
		{
			name:     "custom language and name",
			args:     []string{"cmd", "-lang=fr", "-name=Marie"},
			expected: "Bonjour, Marie![eiraM]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			flag.CommandLine = flag.NewFlagSet(tt.args[0], flag.ExitOnError)
			lang = flag.String("lang", "en", "language")
			name = flag.String("name", "World", "name")

			// Set test args
			os.Args = tt.args

			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)
			actual := buf.String()

			if actual != tt.expected {
				t.Errorf("got %q, want %q", actual, tt.expected)
			}
		})
	}
}

func TestNewGreeter(t *testing.T) {
	greeter := hello.NewGreeter("en", "Test")
	if greeter.Language != "en" {
		t.Errorf("expected language 'en', got %s", greeter.Language)
	}
	if greeter.Name != "Test" {
		t.Errorf("expected name 'Test', got %s", greeter.Name)
	}
}
