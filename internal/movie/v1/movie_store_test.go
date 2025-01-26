package v1

import (
	"phos-metadata-service/internal/movie/v1/models"
	"testing"
)

const testPayloadNotNullMessage = "The payload cannot be null"

func TestCreatedStore(t *testing.T) {
	response := MovieStore().Create(&models.Movie{})

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestUpdateMovieByIdStore(t *testing.T) {
	response := MovieStore().UpdateMovieById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestGetMovieByIdStore(t *testing.T) {
	response := MovieStore().GetMovieById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}

func TestDeleteMovieByIdStore(t *testing.T) {
	response := MovieStore().DeleteMovieById("1")

	if response.Payload == nil {
		t.Error(testPayloadNotNullMessage)
	}
}
