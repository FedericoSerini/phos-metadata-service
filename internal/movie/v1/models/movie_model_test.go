package models

import (
	"encoding/json"
	"testing"
)

var (
	testStatusErrorMessage = "Expected %s but got %s"
	movie1                 = &Movie{
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

	movie2 = &Movie{
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
	}

	movieList = &MovieList{
		ItemCount: 2,
		PageCount: 1,
		Items: []*Movie{
			movie1,
			movie2,
		},
	}

	movieResponse = &MovieResponse{
		Payload:      movie1,
		Status:       200,
		ErrorMessage: nil,
	}

	movieListResponse = &MovieListResponse{
		Payload: []*Movie{
			movie1,
			movie2,
		},
		Status:       200,
		ErrorMessage: nil,
	}

	expectedMovieJson             = `{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"}`
	expectedMovieListJson         = `{"item_count":2,"page_count":1,"items":[{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"},{"media_id":"67890","title":"The Last Hope","original_title":"La Última Esperanza","genre_ids":[18,14,28],"popularity":92.3,"origin_country":["GB","FR"],"vote_count":2500,"release_date":"2025-01-01","backdrop_path":"/images/backdrop_67890.jpg","original_language":"fr","vote_average":7.8,"overview":"A story of redemption and survival.","poster_path":"/images/poster_67890.jpg"}]}`
	expectedMovieResponseJson     = `{"data":{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"}}`
	expectedMovieListResponseJson = `{"data":[{"media_id":"12345","title":"The Great Adventure","original_title":"La Gran Aventura","genre_ids":[28,12,16],"popularity":85.5,"origin_country":["US","ES"],"vote_count":1500,"release_date":"2024-07-15","backdrop_path":"/images/backdrop_12345.jpg","original_language":"en","vote_average":8.2,"overview":"An epic tale of a journey into the unknown.","poster_path":"/images/poster_12345.jpg"},{"media_id":"67890","title":"The Last Hope","original_title":"La Última Esperanza","genre_ids":[18,14,28],"popularity":92.3,"origin_country":["GB","FR"],"vote_count":2500,"release_date":"2025-01-01","backdrop_path":"/images/backdrop_67890.jpg","original_language":"fr","vote_average":7.8,"overview":"A story of redemption and survival.","poster_path":"/images/poster_67890.jpg"}]}`
)

func TestMovieMarshalling(t *testing.T) {
	data, err := json.Marshal(movie1)
	if err != nil {
		t.Fatalf("Error marshalling Movie: %v", err)
	}

	if string(data) != expectedMovieJson {
		t.Errorf(testStatusErrorMessage, expectedMovieJson, string(data))
	}
}

func TestMovieListMarshalling(t *testing.T) {
	data, err := json.Marshal(movieList)
	if err != nil {
		t.Fatalf("Error marshalling MovieList: %v", err)
	}

	if string(data) != expectedMovieListJson {
		t.Errorf(testStatusErrorMessage, expectedMovieListJson, string(data))
	}
}

func TestMovieResponseMarshalling(t *testing.T) {
	data, err := json.Marshal(movieResponse)
	if err != nil {
		t.Fatalf("Error marshalling MovieResponse: %v", err)
	}

	if string(data) != expectedMovieResponseJson {
		t.Errorf(testStatusErrorMessage, expectedMovieResponseJson, string(data))
	}
}

func TestMovieListResponseMarshalling(t *testing.T) {
	data, err := json.Marshal(movieListResponse)
	if err != nil {
		t.Fatalf("Error marshalling MovieListResponse: %v", err)
	}

	if string(data) != expectedMovieListResponseJson {
		t.Errorf(testStatusErrorMessage, expectedMovieListResponseJson, string(data))
	}
}
