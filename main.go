package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exaple.com/note/note"
	"exaple.com/note/todo"
)

type saver interface { // if there's only one method in interface, then name method's name + -er
	Save() error
}

// type displayer interface {
// 	Display() error
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteDate()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) { // or any
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(int)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
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
