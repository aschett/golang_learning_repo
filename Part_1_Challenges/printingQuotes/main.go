package printingquotes

import (
	"bufio"
	"fmt"
	"io"
)

type quoteInformation struct {
	Quote   string
	Speaker string
}

func main() {

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
	return input, nil
}

func createOuputString(quoteInfo quoteInformation) (string, error) {
	if quoteInfo.Quote == "" || quoteInfo.Speaker == "" {
		return "", fmt.Errorf("quoteInformation struct not specified correctly")
	}
	outputString := "\" " + quoteInfo.Quote + " \"" + " - " + quoteInfo.Speaker
	return outputString, nil
}
