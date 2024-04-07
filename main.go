package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
	AGE  int    `json:"age"`
}

var list = []Person{Person{0, "person1", 18}, Person{1, "person2", 19}, Person{0, "person1", 18}, Person{2, "person3", 20}}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // Default port if not specified
	}

	router := gin.Default()
	router.GET("/persons", getPersons)
	router.POST("/post", addPerson)
	router.Run(":" + port)
}

// convert data into json
func getPersons(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, list)
}

func addPerson(context *gin.Context) {
	var pers Person
	var err = context.BindJSON(&pers)
	if err != nil {
		return
	}

	list = append(list, pers)
	fmt.Println(list)
	context.IndentedJSON(http.StatusCreated, pers)
}
