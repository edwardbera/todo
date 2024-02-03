package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type todo struct {
	title       string
	description string
	date        time.Time
	status      bool
}

var collection = make(map[int]todo)

// function to create a todo object
func createTodo() *todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you want to do today?")
	t := todo{}
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
		//key := len(collection) + 1
		//collection[key] = *createTodo()
		insertData(createTodo())
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

	switch option {
	case "1":
		viewTodos()

	case "2":
		inputTodo()
	case "3":
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
	fmt.Println("3. Exit")
	_, err := fmt.Scanln(&reply)
	if err != nil {
		log.Fatal(err)
	}
	return reply
}

// View available tasks
func viewTodos() {
	var id string
	var response string
	var response2 string
	var newstatus string
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}

	sts := "SELECT * FROM todos"

	row, err := db.Query(sts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("-------------------------")
	fmt.Println("Here are your Todos")
	fmt.Println("-------------------------")

	for row.Next() { // Iterate and fetch the records from result cursor
		var title string
		var description string
		var date string
		var status string
		var completed string

		row.Scan(&id, &title, &description, &date, &status)

		if status == "0" {
			completed = "pending"
		} else {
			completed = "completed"
		}

		fmt.Println("ID : "+id, "| Title : "+title, "Description : "+description, "Date : "+strings.Split(date, " ")[0], "Status : "+completed)
		fmt.Println("-------------------------")

	}

	fmt.Println("")
	fmt.Println("--To UPDATE a Todo enter 'U:' followed by the id of the task you wish to update.--")
	fmt.Println("--Enter 'b' to go BACK to menu or 'e' to EXIT the appication.--")

	_, err = fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(response, "U:") {
		var resparr = strings.Split(response, ":")
		fmt.Println("")
		fmt.Println("To delete the todo input D and press Enter")
		fmt.Println("To mark the todo as completed input C and press Enter")
		_, err = fmt.Scanln(&response2)

		if err != nil {
			log.Fatal(err)
		}

		if response2 == "D" {
			deleteTodo(string(resparr[1]))
			fmt.Println("Todo has been deleted.")
			fmt.Println("")
		} else if response2 == "C" {
			newstatus = "1"
			updateTodo(string(resparr[1]), newstatus)
			fmt.Println("Todo status has been updated.")
			fmt.Println("")
			if err != nil {
				log.Fatal(err)
			}
		}

		viewTodos()
	} else {
		switch response {
		case "b":
			option := menu()
			getInput(option)
		case "e":
			break
		default:
			break
		}
		_, err = fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}
		var resparr = strings.Split(response, ":")
		resparr[1] = strings.TrimSpace(string(resparr[1]))

	}

}

// Create SQLite Database
func createDB() {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sts := `
	CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY, title TEXT, description TEXT, date TEXT, status TEXT);
	`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table cars created")
}

func insertData(data *todo) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}

	sts := "INSERT INTO todos(title, description, date, status) VALUES ('" + data.title + "', '" + data.description + "', '" + data.date.String() + "'," + strconv.FormatBool(data.status) + " )"

	_, err = db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func updateTodo(id string, status string) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}
	sts := "UPDATE todos SET status=" + status + " WHERE id=" + id
	_, err = db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func deleteTodo(id string) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}
	sts := "DELETE FROM todos WHERE id=" + id
	_, err = db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	createDB()
	option := menu()
	getInput(option)

}
