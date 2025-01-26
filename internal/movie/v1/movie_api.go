package v1

import (
	"phos-metadata-service/internal/movie/v1/models"

	"github.com/gin-gonic/gin"
)

func AddMovieRoutes(c *gin.RouterGroup) {
	movieRoutes := c.Group("/movie/metadata")
	movieRoutes.POST("", createMovie)
	movieRoutes.PATCH("/:id", updateMovieById)
	movieRoutes.GET("/:id", getMovieById)
	movieRoutes.DELETE("/:id", deleteMovieById)
}

func createMovie(c *gin.Context) {
	var movie models.Movie
	c.BindJSON(&movie)
	res := MovieStore().Create(nil)
	c.JSON(res.Status, res)
}

func updateMovieById(c *gin.Context) {
	id := c.Param("id")
	res := MovieStore().UpdateMovieById(id)
	c.JSON(res.Status, res)
}

func getMovieById(c *gin.Context) {
	id := c.Param("id")
	res := MovieStore().GetMovieById(id)
	c.JSON(res.Status, res)
}

func deleteMovieById(c *gin.Context) {
	id := c.Param("id")
	res := MovieStore().DeleteMovieById(id)
	c.JSON(res.Status, res)
}
