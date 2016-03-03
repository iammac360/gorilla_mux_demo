package main

import (
	"fmt"
	"testing/gorilla_mux_demo/models"
)

var currentId int

var todos models.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(models.Todo{Name: "Write presentation"})
	RepoCreateTodo(models.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) models.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	// return empty Todo if not found
	return models.Todo{}
}

func RepoCreateTodo(t models.Todo) models.Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func ReposDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
