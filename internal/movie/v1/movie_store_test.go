package v1

import (
	"phos-metadata-service/internal/movie/v1/models"
	"testing"
)

func TestCreated(t *testing.T) {
	response := MovieStore().Create(&models.Movie{})

	if response.Payload == nil {
		t.Error("The payload cannot be null")
	}
}

func TestUpdateMovieById(t *testing.T) {
	response := MovieStore().UpdateMovieById("1")

	if response.Payload == nil {
		t.Error("The payload cannot be null")
	}
}

func TestGetMovieById(t *testing.T) {
	response := MovieStore().GetMovieById("1")

	if response.Payload == nil {
		t.Error("The payload cannot be null")
	}
}

func TestDeleteMovieById(t *testing.T) {
	response := MovieStore().DeleteMovieById("1")

	if response.Payload == nil {
		t.Error("The payload cannot be null")
	}
}
