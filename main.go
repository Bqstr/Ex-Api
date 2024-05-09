package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Person struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	MONEY int    `json:"money"`
}

var list = []Person{Person{0, "person1", 1111}, Person{1, "person2", 2222}, Person{0, "person1", 2222}, Person{2, "person3", 2222}}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // Default port if not specified
	}

	router := gin.Default()
	router.GET("/persons", getPersons)
	router.POST("/post", addPerson)
	router.POST("/login", login)
	router.POST("/addmoney", addMoney)

	router.GET("/hey", getRoute)
	router.Run(":" + port)
}

// convert data into json
func getPersons(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, list)
}

func personExists(id int) bool {
	for _, p := range list {
		if p.ID == id {
			return true
		}
	}
	return false
}

func addPerson(context *gin.Context) {
	var requestData struct {
		ID         int `json:"id"`
		MoneyToAdd int `json:"money_to_add"`
	}

	if err := context.BindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	found := false
	for i := range list {
		if list[i].ID == requestData.ID {
			list[i].MONEY += requestData.MoneyToAdd
			context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Added %d money to person with ID %d", requestData.MoneyToAdd, requestData.ID)})
			found = true
			break
		}
	}

	if !found {
		context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Person with ID %d not found", requestData.ID)})
	}
}

func addMoney(context *gin.Context) {
	var requestData struct {
		ID         int `json:"id"`
		MoneyToAdd int `json:"money_to_add"`
	}

	if err := context.BindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	found := false
	for i := range list {
		if list[i].ID == requestData.ID {
			list[i].MONEY += requestData.MoneyToAdd
			context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Added %d money to person with ID %d", requestData.MoneyToAdd, requestData.ID)})
			found = true
			break
		}
	}

	if !found {
		context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Person with ID %d not found", requestData.ID)})
	}
}

func login(context *gin.Context) {
	var pers Person
	var err = context.BindJSON(&pers)
	if err != nil {
		return
	}

	if personExists(pers.ID) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Person with this ID already exists"})
		return
	}

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
