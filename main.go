package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var todos = make(map[int]Todo)

func NewTodos(id int, name string, description string, done bool) Todo {
	return Todo{
		Id:          id,
		Name:        name,
		Description: description,
		Done:        done,
	}
}

func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func PostHandler(c echo.Context) error {
	var todo Todo
	err := c.Bind(&todo)
	todos[todo.Id] = todo
	if err != nil {
		c.String(http.StatusBadRequest, "bad req")
	}
	return c.String(http.StatusOK, "success")
}
func DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	newid, _ := strconv.Atoi(id)

	delete(todos, newid)
	return c.String(http.StatusOK, "deleted")
}

func main() {
	e := echo.New()

	todos[1] = Todo{Id: 1, Name: "Working", Description: "you must do it", Done: false}
	todos[2] = Todo{Id: 2, Name: "Cooking", Description: "you must do it", Done: true}
	todos[3] = Todo{Id: 3, Name: "Cleaning", Description: "you must do it", Done: false}

	e.GET("/todo", GetHandler)
	e.POST("/todo", PostHandler)
	e.DELETE("/todo/{id}", DeleteHandler)
	e.Start(":8080")

}
