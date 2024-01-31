package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type todo struct {
	id          uuid.UUID
	title       string
	description string
	date        time.Time
}

var collection = make(map[int]todo)

// function to create a todo object
func createTodo() *todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you want to do today?")
	t := todo{}
	t.id = uuid.New()
	n, err := reader.ReadString('\n')
	t.title = n
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tell me a little more about your activity")
	d, err2 := reader.ReadString('\n')
	if err2 != nil {
		log.Fatal(err2)
	}
	t.description = d
	t.date = time.Now()
	return &t
}

// function to create multiple todos
func inputTodo() {
	var status string
	for status != "n" {
		key := len(collection) + 1
		collection[key] = *createTodo()
		fmt.Println("Do you want to add another todo? (Enter 'y' for yes or 'n' for no)")
		_, err := fmt.Scanln(&status)
		if err != nil {
			log.Fatal(err)
		}
	}
	option := menu()
	getInput(option)
}

// Receive input from user
func getInput(option string) {
	var response string
	switch option {
	case "1":
		viewTodos()
		fmt.Println("Enter 'b' to go back to menu or 'e' to exit the apllication.")
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}
		switch response {
		case "b":
			option := menu()
			getInput(option)
		case "e":
			break
		default:
			break
		}
	case "2":
		inputTodo()
	case "3":
		deleteTodo()
	case "4":
		break
	default:
		break
	}

}

// Application menu
func menu() string {
	var reply string
	fmt.Println("Menu")
	fmt.Println("1. View Todos")
	fmt.Println("2. Create Todo")
	fmt.Println("3. Delete Todo")
	fmt.Println("4. Exit")
	_, err := fmt.Scanln(&reply)
	if err != nil {
		log.Fatal(err)
	}
	return reply
}

// View available tasks
func viewTodos() {

	fmt.Println("Here are your Todos")
	fmt.Println("Index	  Title")
	fmt.Println("-------------------------")
	for key, value := range collection {
		fmt.Println(key, " | ", value.title)
	}
	fmt.Println("-------------------------")

}

func deleteTodo() {
	var key int
	var response string
	viewTodos()
	fmt.Println("Enter the number for the task you want to delete")
	_, err := fmt.Scan(&key)

	if err != nil {
		log.Fatal(err)
	}

	delete(collection, key)

	fmt.Println("-------------------------")
	fmt.Println("Enter 'b' to go back to menu or 'e' to exit the apllication.")

	_, err2 := fmt.Scanln(&response)
	if err2 != nil {
		log.Fatal(err2)
	}
	switch response {
	case "b":
		option := menu()
		getInput(option)
	case "e":
		break
	default:
		break
	}
}

func main() {
	option := menu()
	getInput(option)
}
