package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	route := gin.Default()

	routeGroup := route.Group("/v1")
	{
		//routeGroup.GET("/user", Hello)
		//routeGroup.GET("/user/:id", Hello)
		routeGroup.POST("/user", Hello)
	}

	err := route.Run(":8000")
	if err != nil {
		panic(err)
	}
}

func Hello(context *gin.Context) {

	//get params
	//id, _ := strconv.Atoi(context.Param("id"))

	//get query string
	//firstName := context.DefaultQuery("first_name", "")
	//lastName := context.DefaultQuery("last_name", "")
	//id, _ := strconv.ParseInt(context.DefaultQuery("id", "0"), 10, 0)

	//get post data
	var personal Personal
	if err := context.BindJSON(&personal); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, personal)
}

type Personal struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}
