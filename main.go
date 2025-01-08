package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func getTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskFromDB(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func getTasks(c *gin.Context) {
	tasks, err := getTasksFromDB()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Problem with getting tasks from collection"})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	var incomingTask struct {
		Title  string `json:"title"`
		IsDone bool   `json:"isDone"`
	}

	if err := c.Bind(&incomingTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task format"})
		return
	}

	task := Task{
		Id:     id,
		Title:  incomingTask.Title,
		IsDone: incomingTask.IsDone,
	}

	foundTask, err := getTaskFromDB(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Problem with updating task in collection"})
		log.Fatal(err)
		return
	}

	*foundTask = task

	err = updateTaskInDB(&task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data"})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func createTask(c *gin.Context) {
	var incomingTask struct {
		Title  string `json:"title"`
		IsDone bool   `json:"isDone"`
	}

	if err := c.Bind(&incomingTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect task structure"})
		return
	}

	if len(incomingTask.Title) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Empty task title"})
		return
	}

	task := Task{
		Id:     uuid.NewString(),
		Title:  incomingTask.Title,
		IsDone: incomingTask.IsDone,
	}

	err := addTaskToDB(&task)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Problem with creating task in database."})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func main() {
	initDb()
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskById)
	router.POST("/createTask", createTask)
	router.PATCH("/updateTask/:id", updateTask)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Panicln(err)
	}
}
