package repository

import (
	"context"
	"phos-metadata-service/internal/tvshow/v1/models"
	"phos-metadata-service/pkg/database"
	"phos-metadata-service/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TvShowRepositoryInterface struct {
	ctx        context.Context
	Collection *mongo.Collection
}

func TvShowRepository() *TvShowRepositoryInterface {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongo.Connect(ctx)
	return &TvShowRepositoryInterface{
		ctx:        ctx,
		Collection: database.GetCollection("test", "tv_shows"),
	}
}

func (repo *TvShowRepositoryInterface) CreateTvShow(tvShow *models.TvShow) (*models.TvShow, error) {
	_, err := repo.Collection.InsertOne(repo.ctx, tvShow)
	if err != nil {
		return nil, err
	}

	return tvShow, nil
}

func (repo *TvShowRepositoryInterface) UpdateTvShowById(id string, tvShow *models.TvShow) (*models.TvShow, error) {
	var replacedTvShow models.TvShow
	err := repo.Collection.FindOneAndReplace(repo.ctx, bson.M{"media_id": tvShow.MediaId}, tvShow).Decode(&replacedTvShow)
	return handleResponse(&replacedTvShow, err)
}

func (repo *TvShowRepositoryInterface) GetTvShowById(id string) (*models.TvShow, error) {
	var tvShow models.TvShow
	err := repo.Collection.FindOne(repo.ctx, bson.M{"media_id": id}).Decode(&tvShow)
	return handleResponse(&tvShow, err)
}

func (repo *TvShowRepositoryInterface) DeleteTvShowById(id string) (*models.TvShow, error) {
	var tvShow models.TvShow
	err := repo.Collection.FindOneAndDelete(repo.ctx, bson.M{"media_id": id}).Decode(&tvShow)
	return handleResponse(&tvShow, err)
}

func handleResponse(movie *models.TvShow, err error) (*models.TvShow, error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.ErrorResourceNotFound
		}
		println(err.Error())
		return nil, utils.ErrorInternal
	}
	return movie, nil
}
