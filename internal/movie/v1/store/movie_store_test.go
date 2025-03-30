package store

import (
	"log"
	"phos-metadata-service/internal/movie/v1/models"
	"phos-metadata-service/internal/movie/v1/repository"
	"phos-metadata-service/pkg/utils"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

const testPayloadNotNullMessage = "The payload cannot be null"

var mongoClient *mongo.Client

// Test Main Setup
func TestMain(m *testing.M) {
	utils.SetupTestDB()
	mongoClient = utils.GetMongoClient()
	_ = m.Run()
	utils.TeardownTestDB()
	log.Println("Test execution finished")
}

// NewMovieRepository for Testing
func newMovieRepository() *repository.MovieRepositoryInterface {
	return &repository.MovieRepositoryInterface{
		Collection: mongoClient.Database("test").Collection("movies"),
	}
}

func TestCreatedStore(t *testing.T) {
	response := MovieStore(newMovieRepository()).Create(&models.Movie{MediaId: "1"})

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestUpdateMovieByIdStore(t *testing.T) {
	response := MovieStore(newMovieRepository()).UpdateMovieById("1", &models.Movie{MediaId: "1", Title: "Inception"})

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestGetMovieByIdStore(t *testing.T) {
	movieRepository := newMovieRepository()
	MovieStore(movieRepository).Create(&models.Movie{MediaId: "2", Title: "Interstellar"})
	response := MovieStore(movieRepository).GetMovieById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestDeleteMovieByIdStore(t *testing.T) {
	response := MovieStore(newMovieRepository()).DeleteMovieById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}
