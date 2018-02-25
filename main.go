package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Completed bool
}

var todosCollection *mgo.Collection
var session *mgo.Session

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	// get a Collection of todo
	todosCollection = session.DB("test-todo").C("todo")
}

func main() {
	defer session.Close()
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}

	router.Run()
}

func createTodo(context *gin.Context) {
	title := context.PostForm("Title")
	completed, _ := strconv.ParseBool(context.PostForm("Completed"))
	var todo = Todo{bson.NewObjectId(), title, completed}
	fmt.Println("" + todo.Title + " completed: " + strconv.FormatBool(todo.Completed))
	err := todosCollection.Insert(&todo)
	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "todo item created successfully",
	})
}

func fetchAllTodo(context *gin.Context) {
	var todos []Todo
	err := todosCollection.Find(nil).All(&todos)
	if err != nil {
		log.Fatal(err)
	}

	if len(todos) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "no todo found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todos,
	})
}

func fetchSingleTodo(context *gin.Context) {
	todo := Todo{}
	id := bson.ObjectIdHex(context.Param("id"))
	err := todosCollection.FindId(id).One(&todo)

	if err != nil || todo == (Todo{}) {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todo,
	})
}

func updateTodo(context *gin.Context) {
	id := bson.ObjectIdHex(context.Param("id"))
	title := context.PostForm("title")
	completed, _ := strconv.ParseBool(context.PostForm("completed"))

	err := todosCollection.UpdateId(id, bson.M{"title": title, "completed": completed})

	fmt.Printf("completed: %t\n\n", completed)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
}

func deleteTodo(context *gin.Context) {
	id := bson.ObjectIdHex(context.Param("id"))

	fmt.Printf("id: %v", id)

	err := todosCollection.RemoveId(id)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully!",
	})
}
