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

func TestReadInput_ValidInput(t *testing.T) {
	defer suppressOutput(t)()
	input := "Test Quote\n"
	reader := strings.NewReader(input)
	result, err := readInput(reader, "")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "Test Quote"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestReadInput_EmptyInput(t *testing.T) {
	defer suppressOutput(t)()
	input := "\n"
	reader := strings.NewReader(input)
	_, err := readInput(reader, "")
	if err == nil {
		t.Fatal("expected error for empty input, got nil")
	}
}

func TestCreateOutputString_Valid(t *testing.T) {
	info := quoteInformation{
		Quote:   "Those are not the droids you are looking for",
		Speaker: "Obi Wan Kenobi",
	}
	result, err := createOuputString(info)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "\"Those are not the droids you are looking for\" - Obi Wan Kenobi\n"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestCreateOutputString_EmptyQuote(t *testing.T) {
	info := quoteInformation{
		Quote:   "",
		Speaker: "Darth Vader",
	}
	_, err := createOuputString(info)
	if err == nil {
		t.Fatal("expected error for empty quote, got nil")
	}
}

func TestCreateOutputString_EmptySpeaker(t *testing.T) {
	info := quoteInformation{
		Quote:   "Luke, I am your Father",
		Speaker: "",
	}
	_, err := createOuputString(info)
	if err == nil {
		t.Fatal("expected error for empty speaker, got nil")
	}
}
