package v1

import "phos-metadata-service/internal/movie/v1/models"

type MovieStoreInterface interface {
	Create() models.MovieResponse
	UpdateMovieById(id string) models.MovieResponse
	GetMovieById(id string) models.MovieResponse
	DeleteMovieById(id string) models.MovieResponse
}

type ctx struct{}

func MovieStore() *ctx {
	return &ctx{}
}

func (ic ctx) Create(movie *models.Movie) models.MovieResponse {
	return models.MovieResponse{Payload: movie, Status: 200}
}

func (ic ctx) UpdateMovieById(id string) models.MovieResponse {
	return models.MovieResponse{Payload: &models.Movie{}, Status: 200}
}

func (ic ctx) GetMovieById(id string) models.MovieResponse {
	return models.MovieResponse{Payload: &models.Movie{}, Status: 200}
}

func (ic ctx) DeleteMovieById(id string) models.MovieResponse {
	return models.MovieResponse{Payload: &models.Movie{}, Status: 200}
}
