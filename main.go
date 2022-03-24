package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed;default=false"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	router.Headers("X-Requested-With", "XMLHttpRequest")

	router.HandleFunc("/", GetTodos).Methods("GET")
	router.HandleFunc("/", SaveTodoController).Methods("POST")
	router.HandleFunc("/{id}", GetTodoController).Methods("GET")
	router.HandleFunc("/{id}", UpdateTodoController).Methods("PUT")
	router.HandleFunc("/{id}", DeleteTodoController).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetTodos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("X-Powered-By", "Yzee")
	json.NewEncoder(res).Encode(todos)
}

func GetTodoController(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, _ := params["id"]

	todo := GetTodo(id)

	if todo.ID == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(todo)
}

func SaveTodoController(res http.ResponseWriter, req *http.Request) {
	var todo Todo
	_ = json.NewDecoder(req.Body).Decode(&todo)

	todo.ID = fmt.Sprintf("%v", len(todos)+1)

	if todo.Title == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Title is required")
		return
	}

	todos = append(todos, todo)

	res.WriteHeader(http.StatusCreated)
}

func UpdateTodoController(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	todo := GetTodo(params["id"])

	if todo.ID == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	var newTodo Todo
	_ = json.NewDecoder(req.Body).Decode(&newTodo)

	todo.Title = newTodo.Title
	todo.Completed = newTodo.Completed

	for index, t := range todos {
		if t.ID == todo.ID {
			todos[index].Completed = newTodo.Completed
			todos[index].Title = newTodo.Title
			break
		}
	}
	fmt.Println(todos)

	json.NewEncoder(res).Encode(todo)
}

func DeleteTodoController(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	todo := GetTodo(params["id"])

	if todo.ID == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	for index, t := range todos {
		if t.ID == todo.ID {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}

	res.WriteHeader(http.StatusNoContent)
}

func GetTodo(id string) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}
