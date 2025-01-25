package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"phos-metadata-service/internal/movie/v1/models"
	"testing"
)

func TestServerRun(t *testing.T) {
	router := getRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Code < 100 {
		t.Errorf(fmt.Sprintf("Response should have a valid http status code, obtained %d", w.Code))
	}
}

func TestCreateMovie(t *testing.T) {
	movieResponse := &models.MovieResponse{}
	router := getRoutes()

	w := httptest.NewRecorder()
	test_payload := "{\"id\":\"12345\",\"media_id\":\"67890\",\"title\":\"The Great Adventure\",\"original_title\":\"La Gran Aventura\",\"genre_ids\":[28,12,16],\"popularity\":42.67,\"origin_country\":[\"US\",\"ES\"],\"vote_count\":1289,\"release_date\":\"2024-07-15\",\"backdrop_path\":\"/images/backdrop_12345.jpg\",\"original_language\":\"en\",\"vote_average\":7.9,\"overview\":\"A thrilling adventure that takes you on a journey through uncharted lands, filled with danger and excitement.\",\"poster_path\":\"/images/poster_12345.jpg\"}"
	req, _ := http.NewRequest("POST", "/v1/movie/metadata", bytes.NewBuffer([]byte(test_payload)))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(fmt.Sprintf("Expected status %d obtained %d", http.StatusOK, w.Code))
	}

	err := json.Unmarshal(w.Body.Bytes(), movieResponse)
	if err != nil {
		t.Error("The response body must be unmarshalled in MovieResponse struct")
	}

	if movieResponse.Payload != nil {
		t.Error("The payload must not be null")
	}
}

func TestGetMovieById(t *testing.T) {
	movieResponse := &models.MovieResponse{}
	router := getRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/movie/metadata/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(fmt.Sprintf("Expected status %d obtained %d", http.StatusOK, w.Code))
	}

	err := json.Unmarshal(w.Body.Bytes(), movieResponse)
	if err != nil {
		t.Error("The response body must be unmarshalled in MovieResponse struct")
	}
}
