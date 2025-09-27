package store

import (
	"log"
	"phos-metadata-service/internal/tvshow/v1/models"
	"phos-metadata-service/internal/tvshow/v1/repository"
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

// NewTvShowRepository for Testing
func newTvShowRepository() *repository.TvShowRepositoryInterface {
	return &repository.TvShowRepositoryInterface{
		Collection: mongoClient.Database("test").Collection("tv_shows"),
	}
}

func TestCreatedStore(t *testing.T) {
	response := TvShowStore(newTvShowRepository()).Create(&models.TvShow{MediaId: "1"})

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestUpdateTvShowByIdStore(t *testing.T) {
	response := TvShowStore(newTvShowRepository()).UpdateTvShowById("1", &models.TvShow{MediaId: "1", Name: "Breaking Bad"})

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestGetTvShowByIdStore(t *testing.T) {
	tvShowRepository := newTvShowRepository()
	TvShowStore(tvShowRepository).Create(&models.TvShow{MediaId: "2", Name: "Berserk"})
	response := TvShowStore(tvShowRepository).GetTvShowById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestDeleteTvShowByIdStore(t *testing.T) {
	response := TvShowStore(newTvShowRepository()).DeleteTvShowById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}
