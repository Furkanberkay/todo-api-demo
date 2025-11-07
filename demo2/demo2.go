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
	return nil
}

func Demo2() {
	e := echo.New()
	e.GET("/demo", getTodos)
	e.POST("/demo", addTodo)
	e.DELETE("/demo/:id", deleteTodos)
	e.PUT("/demo/:id", putTodos)
	e.PATCH("/demo/:id", patchTodos)
	e.Logger.Fatal(e.Start(":8081"))

}
