package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate     time.Time `josn:"due_date"`
	Status       string   `json:"status"`
}

//Mock date for task struct

var tasks = []Task{
	{ID: 1, Title :"Task 1",Description :"First task",DueDate: time.Now(),Status: "done"},
	{ID: 2 , Title :"Task 2",Description :"second task",DueDate: time.Now().AddDate(0,0,2),Status: "ongoing"},
	{ID: 3 , Title :"Task 3",Description :"thied task",DueDate: time.Now().AddDate(0,0,3),Status: "painding"},
}

func gettasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,gin.H{"Tasks": tasks})
}

func gettaskByID( c *gin.Context){

	taskid := c.Param("id")
	c.JSON(http.StatusOK,gin.H{"taskid":taskid})
	Taskid,err  := strconv.Atoi(taskid)
	if err !=  nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"unable to convet to int"})
		return
	}

	for _,task := range tasks{
		if task.ID == Taskid {
			c.IndentedJSON(http.StatusOK,gin.H{"task":task})
			return

		}
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"there is no task with given ID" })
}

func updatetask(c *gin.Context) {
	id := c.Param("id")
	Taskid, err := strconv.Atoi(id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"unable to convert "})
		return 
	}
	var updatedTask Task

	error := c.ShouldBindJSON(&updatedTask)
	if error != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid paylode"})
	}
	for _,task := range tasks{
		if task.ID == Taskid{
			// update only the specified fields
			if updatedTask.Title !=  "" {
				task.Title = updatedTask.Title
			}

			if updatedTask.Description != ""{
				task.Description = updatedTask.Description
			}
			c.IndentedJSON(http.StatusOK,gin.H{"message": "Task updated"})
			return 
		} 

		c.JSON(http.StatusNotFound,gin.H{"message": "Task not found"})
	}
}

func deleteTask(c *gin.Context){
	id := c.Param("id")

	deleteTaskID,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"massage":"can't given id to int"})
	}

	for _,task := range tasks{
		if task.ID == deleteTaskID {
			tasks = append(tasks[:deleteTaskID],tasks[deleteTaskID+1:]... )
			c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted successfuly"})
			return
		}
	}

}

func createTask( c *gin.Context){
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated,gin.H{"message":"Task created"})
	
}
func main() {
	router := gin.Default()
	router.GET("/tasks",gettasks)
	router.GET("/tasks/:id",gettaskByID)
	router.PUT("/tasks/:id",updatetask)
	router.DELETE("/tasks/:id",deleteTask)
	router.POST("/tasks",createTask)

	router.Run("localhost:8080")
}