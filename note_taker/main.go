package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dreking/note-taker/note"
	"github.com/dreking/note-taker/todo"
)

type saver interface {
	Save() error
	Display()
}

type displayer interface {
	Display()
}

// Embedded interface
type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		panic(err)
	}

	err = outputData(todo)
	if err != nil {
		panic(err)
	}

	userNote, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	err = outputData(userNote)
	if err != nil {
		panic(err)
	}

	fmt.Println("saved file")

	printSomething(1)
	printSomething(1.5)
	printSomething("Go")
}

// any data type
func printSomething(value any) {
	// swtich case for checking data type when using any as data type
	switch value.(type) {
	case int:
		fmt.Println("Int", value)
	case float64:
		fmt.Println("Float", value)
	case string:
		fmt.Println("String", value)
	}

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Int", intVal)
	}

	floatVal, ok := value.(int)
	if ok {
		fmt.Println("Float", floatVal)
	}

	stringVal, ok := value.(int)
	if ok {
		fmt.Println("String", stringVal)
	}

	fmt.Println(value)

	add(1, 2)
}

// Generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// var value string
	// This gets only one word with no space
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
