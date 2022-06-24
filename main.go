package main 

import (
	"os"
	"github.com/sunilRF9/go-crud/routes"
	"github.com/gin-gonic/gin"

)


func main () {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// Routes

	router.POST("/entry/create", routes.AddEntry)

	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntryByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)

	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":"+port)
}
