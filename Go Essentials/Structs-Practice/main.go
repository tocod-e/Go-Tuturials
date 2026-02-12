package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface{
	Save() error
}
type outable interface{
	saver
	Display()
}

func main() {
	// Note
	title, content := getNoteDate()
	userNote, err := note.New(title, content)
	if err != nil{
		fmt.Println(err)
		return
	}
	err =outputData(userNote)
	if err != nil {
		return
	}
	
	// Todo
	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil{
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}
}

// The `outputData` function displays data and saves it.
func outputData(data outable) error{
	data.Display()
	return saveData(data)
}

// The `getNoteDate` function in Go prompts the user for a note title and content, and returns them as
// strings.
func getNoteDate() (string, string){
	title  := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

// The `getTodoData` function in Go prompts the user for a todo text input and returns the input as a
// string.
func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

// The `saveData` function in Go attempts to save data using the `Save` method of a `saver` interface
// and prints success or failure messages accordingly.
func saveData(data saver) error{
	err := data.Save()
	if err != nil{
		fmt.Println("Saving the data failed.")
		return err
	} 
	fmt.Println("Saving the data succeeded!.")
	return nil
}

// The function `getUserInput` reads user input from the console in Go, handling cases where the input
// is empty or just a newline.
func getUserInput(prompt string) string {
    fmt.Printf("%v ", prompt)
    // create the reader
    reader := bufio.NewReader(os.Stdin)
    // READ the string
    text, err := reader.ReadString('\n')
    // If the string is empty or just a newline (leftover from previous),
    // try reading ONE more time.
    if strings.TrimSpace(text) == "" && err == nil {
        text, err = reader.ReadString('\n')
    }
    if err != nil {
        return ""
    }
    return strings.TrimSpace(text)
}