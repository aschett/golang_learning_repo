package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type quoteInformation struct {
	Quote   string
	Speaker string
}

func main() {
	var quoteInfo quoteInformation
	var err error
	quoteInfo.Quote, err = readInput(os.Stdin, "What is the Quote?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	quoteInfo.Speaker, err = readInput(os.Stdin, "Who said it?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	outputString, err := createOuputString(quoteInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Print(outputString)
}

func readInput(reader io.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	bufReader := bufio.NewReader(reader)
	input, err := bufReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	if input == "" {
		return "", fmt.Errorf("user did not enter string")
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func createOuputString(quoteInfo quoteInformation) (string, error) {
	if quoteInfo.Quote == "" || quoteInfo.Speaker == "" {
		return "", fmt.Errorf("quoteInformation struct not specified correctly")
	}
	outputString := "\"" + quoteInfo.Quote + "\"" + " - " + quoteInfo.Speaker + "\n"
	return outputString, nil
}
