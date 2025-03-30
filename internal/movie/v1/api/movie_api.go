package api

import (
	"phos-metadata-service/internal/movie/v1/models"
	"phos-metadata-service/internal/movie/v1/repository"
	"phos-metadata-service/internal/movie/v1/store"

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
	res := store.MovieStore(repository.MovieRepository()).Create(&movie)
	c.JSON(res.Status, res)
}

func updateMovieById(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	c.BindJSON(&movie)
	res := store.MovieStore(repository.MovieRepository()).UpdateMovieById(id, &movie)
	c.JSON(res.Status, res)
}

func getMovieById(c *gin.Context) {
	id := c.Param("id")
	res := store.MovieStore(repository.MovieRepository()).GetMovieById(id)
	c.JSON(res.Status, res)
}

func deleteMovieById(c *gin.Context) {
	id := c.Param("id")
	res := store.MovieStore(repository.MovieRepository()).DeleteMovieById(id)
	c.JSON(res.Status, res)
}
