package main

import (
	movieRoutes "phos-metadata-service/internal/movie/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	getRoutes().Run(":9999")
}

func getRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	movieRoutes.AddMovieRoutes(v1)
	return router
}
