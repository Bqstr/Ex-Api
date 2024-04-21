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
	router.GET("/hey", getRoute)
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

func getHey(context_gin *gin.Context) {
	//// Get latitude and longitude parameters from query string
	//start_latitudeStr :=  context_gin.Query("start_latitude")
	//start_longitudeStr:=  context_gin.Query("start_longitude")
	//end_latitudeStr :=    context_gin.Query("end_latitude")
	//end_longitudeStr :=   context_gin.Query("end_longitude")
	//
	//// Convert latitude and longitude to float64
	//start_latitude, err := strconv.ParseFloat(start_latitudeStr, 64)
	//if err != nil {
	//	context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
	//	return
	//}
	//start_longitude, err := strconv.ParseFloat(start_longitudeStr, 64)
	//if err != nil {
	//	context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
	//	return
	//}
	//
	//end_latitude, err := strconv.ParseFloat(end_latitudeStr, 64)
	//if err != nil {
	//	context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
	//	return
	//}
	//end_longitude, err := strconv.ParseFloat(end_longitudeStr, 64)
	//if err != nil {
	//	context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
	//	return
	//}
	//
	////etRoute(start_latitude,start_longitude ,end_latitude, end_longitude)
	//
	//context.String(http.StatusOK, "hey")
}

func get(latitude float64, longitude float64, latitude2 float64, longitude2 float64) {

}
