package main

import (
	"./todo"
	"fmt"
)

func main() {
	history := []interface{}{
		todo.TodoCreated{
			Title:        "Test Task",
			Descriptions: "Hello, World!",
		},
		todo.TodoTitleModified{
			Title: "Test Task Modified",
		},
		todo.TodoInProgress{},
	}

	aggregate := todo.NewFromEvents(history)

	fmt.Println(aggregate)
}
