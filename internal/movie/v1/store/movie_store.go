package store

import (
	"phos-metadata-service/internal/movie/v1/models"
	"phos-metadata-service/internal/movie/v1/repository"
	"phos-metadata-service/pkg/utils"
)

type MovieStoreInterface interface {
	Create() models.MovieResponse
	UpdateMovieById(id string) models.MovieResponse
	GetMovieById(id string) models.MovieResponse
	DeleteMovieById(id string) models.MovieResponse
}

type ctx struct {
	repo *repository.MovieRepositoryInterface
}

func MovieStore(repo *repository.MovieRepositoryInterface) *ctx {
	return &ctx{repo}
}

func (ic ctx) Create(movie *models.Movie) models.MovieResponse {
	res, err := ic.repo.CreateMovie(movie)
	return models.MovieResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) UpdateMovieById(id string, movie *models.Movie) models.MovieResponse {
	res, err := ic.repo.UpdateMovieById(id, movie)
	return models.MovieResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) GetMovieById(id string) models.MovieResponse {
	res, err := ic.repo.GetMovieById(id)
	return models.MovieResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) DeleteMovieById(id string) models.MovieResponse {
	res, err := ic.repo.DeleteMovieById(id)
	return models.MovieResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}
