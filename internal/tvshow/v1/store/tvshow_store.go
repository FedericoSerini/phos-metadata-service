package store

import (
	"phos-metadata-service/internal/tvshow/v1/models"
	"phos-metadata-service/internal/tvshow/v1/repository"
	"phos-metadata-service/pkg/utils"
)

type TvShowStoreInterface interface {
	Create() models.TvShowResponse
	UpdateTvShowById(id string) models.TvShowResponse
	GetTvShowById(id string) models.TvShowResponse
	DeleteTvShowById(id string) models.TvShowResponse
}

type ctx struct {
	repo *repository.TvShowRepositoryInterface
}

func TvShowStore(repo *repository.TvShowRepositoryInterface) *ctx {
	return &ctx{repo}
}

func (ic ctx) Create(tvShow *models.TvShow) models.TvShowResponse {
	res, err := ic.repo.CreateTvShow(tvShow)
	return models.TvShowResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) UpdateTvShowById(id string, tvShow *models.TvShow) models.TvShowResponse {
	res, err := ic.repo.UpdateTvShowById(id, tvShow)
	return models.TvShowResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) GetTvShowById(id string) models.TvShowResponse {
	res, err := ic.repo.GetTvShowById(id)
	return models.TvShowResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}

func (ic ctx) DeleteTvShowById(id string) models.TvShowResponse {
	res, err := ic.repo.DeleteTvShowById(id)
	return models.TvShowResponse{Payload: res, Status: utils.HandleWebStatusErrors(err)}
}
