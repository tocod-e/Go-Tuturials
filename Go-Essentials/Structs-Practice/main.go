package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
type outable interface {
	saver
	Display()
}

func main() {
	// Note
	title, content := getNoteDate()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		return
	}

	// Todo
	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}
}

func outputData(data outable) error {
	data.Display()
	return saveData(data)
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}
	fmt.Println("Saving the data succeeded!.")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if strings.TrimSpace(text) == "" && err == nil {
		text, err = reader.ReadString('\n')
	}
	if err != nil {
		return ""
	}
	return strings.TrimSpace(text)
}
