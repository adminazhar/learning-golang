package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int
	Author string
	Item   string
}

var items = []Todo{
	{ID: 1, Author: "Azhar", Item: "Take out this item"},
	{ID: 2, Author: "Azhar", Item: "Buy pipeapples"},
	{ID: 3, Author: "Batman", Item: "Take batmobile to service"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func main() {

	for _, value := range items {
		fmt.Println(value)
	}

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
}
