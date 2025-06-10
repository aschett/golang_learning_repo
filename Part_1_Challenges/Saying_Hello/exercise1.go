package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
	"log"
	"strings"
)

// I am well aware this may be unefficient but the Book set those constraints

func main(){
	name := readInput(os.Stdin)
	greeting := buildGreetingString(name)
	printGreeting(greeting)
}


func readInput(reader io.Reader) (name string){
	fmt.Println("What is your name?")
	bufReader := bufio.NewReader(reader)
	name, err := bufReader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil{
		log.Fatal(err)
	}
	if name == "" {
		log.Fatal("You did not enter a name.")
	}
	return name
}

func buildGreetingString(name string) string {
	var builder strings.Builder
	builder.WriteString("Hello " + name + ", nice to meet you!")
	return builder.String()
}


func printGreeting(greeting string){
	fmt.Println(greeting)
}