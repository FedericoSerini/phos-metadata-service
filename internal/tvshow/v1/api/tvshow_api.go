package api

import (
	"phos-metadata-service/internal/tvshow/v1/models"
	"phos-metadata-service/internal/tvshow/v1/repository"
	"phos-metadata-service/internal/tvshow/v1/store"

	"github.com/gin-gonic/gin"
)

func AddTvShowRoutes(c *gin.RouterGroup) {
	tvShowRoutes := c.Group("/tvshow/metadata")
	tvShowRoutes.POST("", createTvShow)
	tvShowRoutes.PATCH("/:id", updateTvShowById)
	tvShowRoutes.GET("/:id", getTvShowById)
	tvShowRoutes.DELETE("/:id", deleteTvShowById)
}

func createTvShow(c *gin.Context) {
	var tvShow models.TvShow
	c.BindJSON(&tvShow)
	res := store.TvShowStore(repository.TvShowRepository()).Create(&tvShow)
	c.JSON(res.Status, res)
}

func updateTvShowById(c *gin.Context) {
	id := c.Param("id")
	var tvShow models.TvShow
	c.BindJSON(&tvShow)
	res := store.TvShowStore(repository.TvShowRepository()).UpdateTvShowById(id, &tvShow)
	c.JSON(res.Status, res)
}

func getTvShowById(c *gin.Context) {
	id := c.Param("id")
	res := store.TvShowStore(repository.TvShowRepository()).GetTvShowById(id)
	c.JSON(res.Status, res)
}

func deleteTvShowById(c *gin.Context) {
	id := c.Param("id")
	res := store.TvShowStore(repository.TvShowRepository()).DeleteTvShowById(id)
	c.JSON(res.Status, res)
}
