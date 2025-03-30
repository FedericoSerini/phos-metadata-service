package repository

import (
	"context"
	"phos-metadata-service/internal/movie/v1/models"
	"phos-metadata-service/pkg/database"
	"phos-metadata-service/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepositoryInterface struct {
	ctx        context.Context
	Collection *mongo.Collection
}

func MovieRepository() *MovieRepositoryInterface {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongo.Connect(ctx)
	return &MovieRepositoryInterface{
		ctx:        ctx,
		Collection: database.GetCollection("test", "movies"),
	}
}

func (repo *MovieRepositoryInterface) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	_, err := repo.Collection.InsertOne(repo.ctx, movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (repo *MovieRepositoryInterface) UpdateMovieById(id string, movie *models.Movie) (*models.Movie, error) {
	var replacedMovie models.Movie
	err := repo.Collection.FindOneAndReplace(repo.ctx, bson.M{"media_id": movie.MediaId}, movie).Decode(&replacedMovie)
	return handleResponse(&replacedMovie, err)
}

func (repo *MovieRepositoryInterface) GetMovieById(id string) (*models.Movie, error) {
	var movie models.Movie
	err := repo.Collection.FindOne(repo.ctx, bson.M{"media_id": id}).Decode(&movie)
	return handleResponse(&movie, err)
}

func (repo *MovieRepositoryInterface) DeleteMovieById(id string) (*models.Movie, error) {
	var movie models.Movie
	err := repo.Collection.FindOneAndDelete(repo.ctx, bson.M{"media_id": id}).Decode(&movie)
	return handleResponse(&movie, err)
}

func handleResponse(movie *models.Movie, err error) (*models.Movie, error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.ErrorResourceNotFound
		}
		println(err.Error())
		return nil, utils.ErrorInternal
	}
	return movie, nil
}
