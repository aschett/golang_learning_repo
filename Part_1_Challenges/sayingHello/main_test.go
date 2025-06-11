package main

import (
	"os"
	"strings"
	"testing"
)

func suppressOutput(t *testing.T) func() {
	nullFile, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("failed to open os.DevNull: %v", err)
	}
	original := os.Stdout
	os.Stdout = nullFile
	return func() {
		os.Stdout = original
		nullFile.Close()
	}
}

func TestBuildGreetingString(t *testing.T) {
	name := "John"
	expected := "Hello John, nice to meet you!"
	result := buildGreetingString(name)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestReadInput_Valid(t *testing.T) {
	defer suppressOutput(t)()
	input := "Andi\n"
	reader := strings.NewReader(input)
	result, err := readInput(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := "Andi"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestReadInput_Japanese(t *testing.T) {
	defer suppressOutput(t)()
	input := "尾田 栄一郎\n"
	reader := strings.NewReader(input)
	result, err := readInput(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := "尾田 栄一郎"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestReadInput_Emoji(t *testing.T) {
	defer suppressOutput(t)()
	input := "😱\n"
	reader := strings.NewReader(input)
	result, err := readInput(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := "😱"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestReadInput_Empty(t *testing.T) {
	defer suppressOutput(t)()
	input := "\n"
	reader := strings.NewReader(input)
	_, err := readInput(reader)
	if err == nil {
		t.Error("Expected error for empty input, got none")
	}
	if !strings.Contains(err.Error(), "did not enter a name") {
		t.Errorf("Expected 'did not enter a name' error, got: %v", err)
	}
}
