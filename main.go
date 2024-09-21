package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exaple.com/note/note"
)

func main() {
	title, content := getNoteDate()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
