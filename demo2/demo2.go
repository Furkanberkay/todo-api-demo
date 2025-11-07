package demo2

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTodoDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

var todos = make(map[string]*Todo)

func getTodos(c echo.Context) error {
	req := c.Request().Header.Get("Accept")
	fmt.Println(req)

	return c.JSON(http.StatusOK, todos)

}

func addTodo(c echo.Context) error {
	todo := Todo{}
	err := c.Bind(&todo)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if _, ok := todos[todo.Id]; ok {
		return c.JSON(http.StatusConflict, "todo with this id already exists")
	}
	if todo.Name == "" {
		return c.JSON(http.StatusBadRequest, "name is required")
	}
	if todo.Description == "" {
		return c.JSON(http.StatusBadRequest, "Description is required")
	}
	if todo.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	todos[todo.Id] = &todo
	return c.JSON(http.StatusOK, todos[todo.Id])
}

func deleteTodos(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	if _, ok := todos[id]; !ok {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	delete(todos, id)
	return c.JSON(http.StatusOK, "Deleted")
}

func putTodos(c echo.Context) error {
	id := c.Param("id")

	result, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusNotFound, " todo is not found")

	}

	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result.Id = todo.Id
	result.Name = todo.Name
	result.Description = todo.Description
	result.Completed = todo.Completed

	return c.JSON(http.StatusOK, "success")

}

func patchTodos(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is req")
	}
	value, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusBadGateway, "")
	}
	todoDto := UpdateTodoDTO{}

	if err := c.Bind(&todoDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if todoDto.Description != nil || *todoDto.Description != "" {
		value.Description = *todoDto.Description
	}
	if todoDto.Name != nil || *todoDto.Name != "" {
		value.Name = *todoDto.Name
	}
	if todoDto.Completed != nil {
		value.Completed = *todoDto.Completed
	}
	return c.JSON(http.StatusOK, value)
}

func Demo2() {
	e := echo.New()
	e.GET("/demo", getTodos)
	e.POST("/demo", addTodo)
	e.DELETE("/demo/:id", deleteTodos)
	e.PUT("/demo/:id", putTodos)
	e.PATCH("/demo/:id", patchTodos)
	e.Logger.Fatal(e.Start(":8082"))

}
