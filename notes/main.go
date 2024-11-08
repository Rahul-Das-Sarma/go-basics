package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNotes()

	n, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}
	n.Display()
}

func getNotes() (string, string) {
	title := getUserInput("Please enter a title: ")

	content := getUserInput("Please enter a content: ")

	return title, content
}
func getUserInput(print string) string {
	var value string
	fmt.Print(print)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	value = strings.TrimSuffix(text, "\n")

	return value
}
