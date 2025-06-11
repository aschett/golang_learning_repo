package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// I am well aware this may be unefficient but the Book set those constraints

func main() {
	name, err := readInput(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	greeting := buildGreetingString(name)
	printGreeting(greeting)
}

func readInput(reader io.Reader) (string, error) {
	fmt.Println("What is your name?")
	bufReader := bufio.NewReader(reader)
	name, err := bufReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("you did not enter a name")
	}
	return name, nil
}

func buildGreetingString(name string) string {
	var builder strings.Builder
	builder.WriteString("Hello " + name + ", nice to meet you!")
	return builder.String()
}

func printGreeting(greeting string) {
	fmt.Println(greeting)
}
