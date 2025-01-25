package models

import (
	"encoding/json"
	"testing"
)

func TestMovieMarshalling(t *testing.T) {
	movie := &Movie{
		MediaId:          "12345",
		Title:            "The Great Adventure",
		OriginalTitle:    "La Gran Aventura",
		GenreIds:         []int16{28, 12, 16},
		Popularity:       85.5,
		OriginCountry:    []string{"US", "ES"},
		VoteCount:        1500,
		ReleaseDate:      "2024-07-15",
		BackdropPath:     "/images/backdrop_12345.jpg",
		OriginalLanguage: "en",
		VoteAverage:      8.2,
		Overview:         "An epic tale of a journey into the unknown.",
		PosterPath:       "/images/poster_12345.jpg",
	}

	// Test Marshalling
	data, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Error marshalling Movie: %v", err)
	}

	// Check if the JSON contains expected values
	expected := `{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"}`
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}

func TestMovieListMarshalling(t *testing.T) {
	movieList := &MovieList{
		ItemCount: 2,
		PageCount: 1,
		Items: []*Movie{
			{
				MediaId:          "12345",
				Title:            "The Great Adventure",
				OriginalTitle:    "La Gran Aventura",
				GenreIds:         []int16{28, 12, 16},
				Popularity:       85.5,
				OriginCountry:    []string{"US", "ES"},
				VoteCount:        1500,
				ReleaseDate:      "2024-07-15",
				BackdropPath:     "/images/backdrop_12345.jpg",
				OriginalLanguage: "en",
				VoteAverage:      8.2,
				Overview:         "An epic tale of a journey into the unknown.",
				PosterPath:       "/images/poster_12345.jpg",
			},
			{
				MediaId:          "67890",
				Title:            "The Last Hope",
				OriginalTitle:    "La Última Esperanza",
				GenreIds:         []int16{18, 14, 28},
				Popularity:       92.3,
				OriginCountry:    []string{"GB", "FR"},
				VoteCount:        2500,
				ReleaseDate:      "2025-01-01",
				BackdropPath:     "/images/backdrop_67890.jpg",
				OriginalLanguage: "fr",
				VoteAverage:      7.8,
				Overview:         "A story of redemption and survival.",
				PosterPath:       "/images/poster_67890.jpg",
			},
		},
	}

	// Test Marshalling
	data, err := json.Marshal(movieList)
	if err != nil {
		t.Fatalf("Error marshalling MovieList: %v", err)
	}

	// Check if the JSON contains expected values
	expected := `{"item_count":2,"page_count":1,"items":[{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"},{"media_id":"67890","title":"The Last Hope","original_title":"La Última Esperanza","genre_ids":[18,14,28],"popularity":92.3,"origin_country":["GB","FR"],"vote_count":2500,"release_date":"2025-01-01","backdrop_path":"/images/backdrop_67890.jpg","original_language":"fr","vote_average":7.8,"overview":"A story of redemption and survival.","poster_path":"/images/poster_67890.jpg"}]}`
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}

func TestMovieResponseMarshalling(t *testing.T) {
	movieResponse := &MovieResponse{
		Payload: &Movie{
			MediaId:          "12345",
			Title:            "The Great Adventure",
			OriginalTitle:    "La Gran Aventura",
			GenreIds:         []int16{28, 12, 16},
			Popularity:       85.5,
			OriginCountry:    []string{"US", "ES"},
			VoteCount:        1500,
			ReleaseDate:      "2024-07-15",
			BackdropPath:     "/images/backdrop_12345.jpg",
			OriginalLanguage: "en",
			VoteAverage:      8.2,
			Overview:         "An epic tale of a journey into the unknown.",
			PosterPath:       "/images/poster_12345.jpg",
		},
		Status:       200,
		ErrorMessage: nil,
	}

	// Test Marshalling
	data, err := json.Marshal(movieResponse)
	if err != nil {
		t.Fatalf("Error marshalling MovieResponse: %v", err)
	}

	// Check if the JSON contains expected values
	expected := `{"data":{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"}}`
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}

func TestMovieListResponseMarshalling(t *testing.T) {
	movieListResponse := &MovieListResponse{
		Payload: []*Movie{
			{
				MediaId:          "12345",
				Title:            "The Great Adventure",
				OriginalTitle:    "La Gran Aventura",
				GenreIds:         []int16{28, 12, 16},
				Popularity:       85.5,
				OriginCountry:    []string{"US", "ES"},
				VoteCount:        1500,
				ReleaseDate:      "2024-07-15",
				BackdropPath:     "/images/backdrop_12345.jpg",
				OriginalLanguage: "en",
				VoteAverage:      8.2,
				Overview:         "An epic tale of a journey into the unknown.",
				PosterPath:       "/images/poster_12345.jpg",
			},
			{
				MediaId:          "67890",
				Title:            "The Last Hope",
				OriginalTitle:    "La Última Esperanza",
				GenreIds:         []int16{18, 14, 28},
				Popularity:       92.3,
				OriginCountry:    []string{"GB", "FR"},
				VoteCount:        2500,
				ReleaseDate:      "2025-01-01",
				BackdropPath:     "/images/backdrop_67890.jpg",
				OriginalLanguage: "fr",
				VoteAverage:      7.8,
				Overview:         "A story of redemption and survival.",
				PosterPath:       "/images/poster_67890.jpg",
			},
		},
		Status:       200,
		ErrorMessage: nil,
	}

	// Test Marshalling
	data, err := json.Marshal(movieListResponse)
	if err != nil {
		t.Fatalf("Error marshalling MovieListResponse: %v", err)
	}

	// Check if the JSON contains expected values
	expected := `{"data":[{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"},{"media_id":"67890","title":"The Last Hope","original_title":"La Última Esperanza","genre_ids":[18,14,28],"popularity":92.3,"origin_country":["GB","FR"],"vote_count":2500,"release_date":"2025-01-01","backdrop_path":"/images/backdrop_67890.jpg","original_language":"fr","vote_average":7.8,"overview":"A story of redemption and survival.","poster_path":"/images/poster_67890.jpg"}]}`
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}
