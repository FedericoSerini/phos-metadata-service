package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"phos-metadata-service/internal/movie/v1/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	errorMessage    = "Resource not found"
	expectedJson    = `{"data":{"backdrop_path":"", "genre_ids":[], "media_id":"", "origin_country":[], "original_language":"", "original_title":"", "overview":"", "popularity":0, "poster_path":"", "release_date":"", "title":"", "vote_average":0, "vote_count":0}}`
	responsePayload = models.Movie{GenreIds: []int16{}, OriginCountry: []string{}}
)

type MockMovieStore struct{}

func (m *MockMovieStore) Create(_ interface{}) models.MovieResponse {
	// Mock success response for Create
	return models.MovieResponse{
		Status:       http.StatusCreated,
		ErrorMessage: nil,
	}
}

func (m *MockMovieStore) UpdateMovieById(id string) models.MovieResponse {
	// Mock success response for Update
	if id == "valid-id" {
		return models.MovieResponse{
			Payload:      &responsePayload,
			Status:       http.StatusOK,
			ErrorMessage: nil,
		}
	}
	return models.MovieResponse{
		Payload:      nil,
		Status:       http.StatusNotFound,
		ErrorMessage: &errorMessage,
	}
}

func (m *MockMovieStore) GetMovieById(id string) models.MovieResponse {
	// Mock response for GetMovieById
	if id == "valid-id" {
		return models.MovieResponse{
			Status:       http.StatusOK,
			ErrorMessage: nil,
		}
	}
	return models.MovieResponse{
		Status:       http.StatusNotFound,
		ErrorMessage: &errorMessage,
	}
}

func (m *MockMovieStore) DeleteMovieById(id string) models.MovieResponse {
	// Mock response for DeleteMovieById
	if id == "valid-id" {
		return models.MovieResponse{
			Status:       http.StatusOK,
			ErrorMessage: nil,
		}
	}
	return models.MovieResponse{
		Status:       http.StatusNotFound,
		ErrorMessage: &errorMessage,
	}
}

// Helper function to set up routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddMovieRoutes(v1)
	return router
}

func TestCreateMovie(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/v1/movie/metadata", bytes.NewBuffer([]byte(expectedJson)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.JSONEq(t, expectedJson, rr.Body.String())
}

func TestUpdateMovieById(t *testing.T) {
	router := setupRouter()

	// Test valid update
	req, _ := http.NewRequest("PATCH", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.JSONEq(t, expectedJson, rr.Body.String())
}

func TestGetMovieById(t *testing.T) {
	router := setupRouter()

	// Test valid fetch
	req, _ := http.NewRequest("GET", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.JSONEq(t, expectedJson, rr.Body.String())
}

func TestDeleteMovieById(t *testing.T) {
	router := setupRouter()

	// Test valid delete
	req, _ := http.NewRequest("DELETE", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.JSONEq(t, expectedJson, rr.Body.String())
}
