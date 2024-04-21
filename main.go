package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
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
	router.GET("/hey", getHey)
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

func getHey(context *gin.Context) {
	// Get latitude and longitude parameters from query string
	latitudeStr := context.Query("latitude")
	longitudeStr := context.Query("longitude")

	// Convert latitude and longitude to float64
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}
	fmt.Println(longitude, latitude)

	// Your logic here to handle latitude and longitude...
	// For now, simply return the string "hey"
	context.String(http.StatusOK, "hey")
}
