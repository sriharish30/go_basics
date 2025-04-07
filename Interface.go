package main

import (
	"bufio"
	"example/note/note"
	"example/note/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

//	type displayer interface {
//		display()
//		save()
//	}
type outputable interface {
	saver
	Display()
}

func main() {
	printSomething("sri")
	printSomething(1233)
	printSomething(90.7)
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

	outputData(userNote)

}
func SaveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving the Note failed")
		return err

	}
	fmt.Println("saving the Note suceeded")
	return nil

}

//Interface

func outputData(data outputable) error {
	data.Display()
	return SaveData(data)
}

// In this function the type interface can accept all types of data like int string or others the another name of interface is (any)we can use any also func printsomething(text any)like this

func printSomething(text interface{}) {
	//Using switch case in this function to know the type of the given input
	switch text.(type) {
	case int:
		fmt.Println("Integer: ", text)
	case float64:
		fmt.Println("Float: ", text)
	case string:
		fmt.Println("String :", text)
	}

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")

	return text
}
