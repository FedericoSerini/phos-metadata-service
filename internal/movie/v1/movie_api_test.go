package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var inputJson = `{"data":{"backdrop_path":"", "genre_ids":[], "media_id":"", "origin_country":[], "original_language":"", "original_title":"", "overview":"", "popularity":0, "poster_path":"", "release_date":"", "title":"", "vote_average":0, "vote_count":0}}`

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddMovieRoutes(v1)
	return router
}

func TestCreateMovie(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/v1/movie/metadata", bytes.NewBuffer([]byte(inputJson)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateMovieById(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("PATCH", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetMovieById(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteMovieById(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/v1/movie/metadata/valid-id", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
