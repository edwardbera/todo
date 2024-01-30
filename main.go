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

// function to create a todo
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

// Receive input from user
func getInput() {
	var status string

	for status != "exit" {
		key := len(collection)
		collection[key] = *createTodo()
		_, err := fmt.Scanln(&status)
		if err != nil {
			log.Fatal(err)
		}
	}

	for k, v := range collection {
		fmt.Print(k, v.description)
	}
}

func main() {
	getInput()
}
