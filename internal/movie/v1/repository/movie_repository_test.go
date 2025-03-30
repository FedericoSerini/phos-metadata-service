package repository

import (
	"log"
	"phos-metadata-service/internal/movie/v1/models"
	"phos-metadata-service/pkg/utils"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

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
func newMovieRepository() *MovieRepositoryInterface {
	return &MovieRepositoryInterface{
		Collection: mongoClient.Database("test").Collection("movies"),
	}
}

func TestCreateMovie(t *testing.T) {
	repo := newMovieRepository()

	movie := &models.Movie{
		MediaId:  "1",
		Title:    "Inception",
		GenreIds: []int16{12},
	}

	createdMovie, err := repo.CreateMovie(movie)
	if err != nil {
		t.Fatalf("Failed to create movie: %v", err)
	}

	if createdMovie.Title != "Inception" {
		t.Errorf("Expected 'Inception', got %s", createdMovie.Title)
	}
}

func TestGetMovieById(t *testing.T) {
	repo := newMovieRepository()

	// Insert test movie
	movie := &models.Movie{
		MediaId:  "2",
		Title:    "Interstellar",
		GenreIds: []int16{12},
	}
	_, err := repo.CreateMovie(movie)
	if err != nil {
		t.Fatalf("Failed to insert test movie: %v", err)
	}

	// Fetch movie
	foundMovie, err := repo.GetMovieById(movie.MediaId)
	if err != nil {
		t.Fatalf("Failed to get movie by ID: %v", err)
	}

	if foundMovie.Title != "Interstellar" {
		t.Errorf("Expected 'Interstellar', got %s", foundMovie.Title)
	}
}

func TestUpdateMovieById(t *testing.T) {
	repo := newMovieRepository()

	// Insert a test movie
	movie := &models.Movie{
		MediaId:  "3",
		Title:    "The Matrix",
		GenreIds: []int16{12},
	}
	_, err := repo.CreateMovie(movie)
	if err != nil {
		t.Fatalf("Failed to insert test movie: %v", err)
	}

	// Update movie
	updatedMovie := &models.Movie{
		MediaId:  "3",
		Title:    "The Matrix Reloaded",
		GenreIds: []int16{12},
	}

	_, err = repo.UpdateMovieById(movie.MediaId, updatedMovie)
	if err != nil {
		t.Fatalf("Failed to update movie: %v", err)
	}

	// Fetch updated movie
	fetchedMovie, err := repo.GetMovieById(movie.MediaId)
	if err != nil {
		t.Fatalf("Failed to get updated movie: %v", err)
	}

	if fetchedMovie.Title != "The Matrix Reloaded" {
		t.Errorf("Expected 'The Matrix Reloaded', got %s", fetchedMovie.Title)
	}
}

func TestDeleteMovieById(t *testing.T) {
	repo := newMovieRepository()

	// Insert a test movie
	movie := &models.Movie{
		MediaId:  "4",
		Title:    "Avatar",
		GenreIds: []int16{12},
	}
	_, err := repo.CreateMovie(movie)
	if err != nil {
		t.Fatalf("Failed to insert test movie: %v", err)
	}

	// Delete the movie
	_, err = repo.DeleteMovieById(movie.MediaId)
	if err != nil {
		t.Fatalf("Failed to delete movie: %v", err)
	}

	// Try to fetch deleted movie
	_, err = repo.GetMovieById(movie.MediaId)
	if err == nil {
		t.Fatalf("Expected error, got nil (movie should not exist)")
	}
}
