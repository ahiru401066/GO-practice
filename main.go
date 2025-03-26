package main

import "github.com/gin-gonic/gin"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todos = []Todo{
	{ID: 1, Title: "Goの勉強", Description: "Ginフレームワークを学ぶ", Completed: false},
	{ID: 2, Title: "買い物", Description: "牛乳と卵を買う", Completed: false},
}

func main() {
	r := gin.Default()
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.BindJSON(&newTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		// IDを増やす
		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
		c.JSON(201, newTodo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTodo Todo

		if err := c.BindJSON(&updatedTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
	
		for i, todo := range todos {
			if todo.ID == id {
				todos[i] = updatedTodo
				c.JSON(200, updatedTodo)
				return
			}
		}
	
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
	
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(200, gin.H{"message": "Deleted successfully"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	
	r.Run()
}