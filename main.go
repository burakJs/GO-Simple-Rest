package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID         string `json:"id"`
	Todo       string `json:"todo"`
	IsComplete bool   `json:"isComplete"`
}

var todos []Todo = []Todo{
	{ID: "1", Todo: "Todo - 1", IsComplete: false},
	{ID: "2", Todo: "Todo - 2", IsComplete: true},
	{ID: "3", Todo: "Todo - 3", IsComplete: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if error := context.BindJSON(&newTodo); error != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByIdLogic(id string) (*Todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByIdLogic(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.IndentedJSON(http.StatusFound, todo)
	}

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todo/:id", getTodoById)
	router.Run("localhost:8000")
}
