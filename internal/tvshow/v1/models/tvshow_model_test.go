package models

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testStatusErrorMessage = "Expected %s but got %s"
	tvShow1                = &TvShow{
		id:           primitive.NewObjectID(),
		BackdropPath: "/backdrop/fake_show.jpg",
		CreatedBy: []*CreatedBy{
			{Id: 1, CreditId: "cred123", Name: "Jane Doe", Gender: 1, ProfilePath: "/profile/jane.jpg"},
		},
		EpisodeRuntime:   []int32{45},
		FirstAirDate:     "2020-01-01",
		Genres:           []*Genre{{Id: 1, Name: "Drama"}, {Id: 2, Name: "Sci-Fi"}},
		Homepage:         "fakeshow.example.com",
		MediaId:          "media123",
		InProduction:     true,
		Languages:        []string{"en", "fr"},
		LastAirDate:      "2022-12-12",
		LastEpisodeToAir: &LastEpisodeToAir{Id: 101, EpisodeNumber: 10, AirDate: "2022-12-12", Name: "Finale", Overview: "The big ending", ProductionCode: "S02E10", SeasonNumber: 2, ShowId: 1, StillPath: "/still/finale.jpg", VoteAverage: 8.5, VoteCount: 100},
		Name:             "Fake Show",
		NextEpisodeToAir: "",
		Networks:         []*Network{{Id: 1, Name: "Fake Network", LogoPath: "/logos/network.png", OriginCountry: "US"}},
		NumberOfEpisodes: 20,
		NumberOfSeasons:  2,
		OriginCountry:    []string{"US"},
		OriginalLanguage: "en",
		OriginalName:     "Fake Show Original",
		Overview:         "This is a fake show used for testing.",
		Popularity:       123.45,
		PosterPath:       "/poster/fake.jpg",
		ProductionCompanies: []*ProductionCompany{
			{Id: 1, Name: "Fake Studios", LogoPath: "/logos/fakestudios.png", OriginCountry: "US"},
		},
		Seasons: []*Season{
			{Id: 1, AirDate: "2020-01-01", EpisodeCount: 10, Overview: "Season 1 overview", PosterPath: "/poster/season1.jpg", SeasonNumber: 1},
			{Id: 2, AirDate: "2022-01-01", EpisodeCount: 10, Overview: "Season 2 overview", PosterPath: "/poster/season2.jpg", SeasonNumber: 2},
		},
		Episodes: []*Episode{
			{MediaId: "ep1", Title: "Pilot", OriginalTitle: "Pilot Original", Overview: "Episode 1 overview", PosterPath: "/poster/ep1.jpg", EpisodeNumber: 1},
		},
		Status:      "Ended",
		Type:        "Scripted",
		VoteAverage: 8.9,
		VoteCount:   250,
	}

	expectedTvShowJson = `{"backdrop_path":"/backdrop/fake_show.jpg","created_by":[{"id":1,"credit_id":"cred123","name":"Jane Doe","gender":1,"profile_path":"/profile/jane.jpg"}],"episode_run_time":[45],"first_air_date":"2020-01-01","genres":[{"id":1,"name":"Drama"},{"id":2,"name":"Sci-Fi"}],"homepage":"fakeshow.example.com","media_id":"media123","in_production":true,"languages":["en","fr"],"last_air_date":"2022-12-12","last_episode_to_air":{"id":101,"episode_number":10,"air_date":"2022-12-12","name":"Finale","overview":"The big ending","production_code":"S02E10","season_number":2,"show_id":1,"still_path":"/still/finale.jpg","vote_average":8.5,"vote_count":100},"name":"Fake Show","next_episode_to_air":"","networks":[{"id":1,"name":"Fake Network","logo_path":"/logos/network.png","origin_country":"US"}],"number_of_episodes":20,"number_of_seasons":2,"origin_country":["US"],"original_language":"en","original_name":"Fake Show Original","overview":"This is a fake show used for testing.","popularity":123.45,"poster_path":"/poster/fake.jpg","production_companies":[{"id":1,"name":"Fake Studios","logo_path":"/logos/fakestudios.png","origin_country":"US"}],"seasons":[{"id":1,"air_date":"2020-01-01","episode_count":10,"overview":"Season 1 overview","poster_path":"/poster/season1.jpg","season_number":1},{"id":2,"air_date":"2022-01-01","episode_count":10,"overview":"Season 2 overview","poster_path":"/poster/season2.jpg","season_number":2}],"episodes":[{"media_id":"ep1","title":"Pilot","original_title":"Pilot Original","overview":"Episode 1 overview","poster_path":"/poster/ep1.jpg","episode_number":1}],"status":"Ended","type":"Scripted","vote_average":8.9,"vote_count":250}`

	tvShowList = &TvShowList{
		ItemCount: 1,
		PageCount: 1,
		Items:     []*TvShow{tvShow1},
	}

	expectedTvShowListJson = `{"item_count":1,"page_count":1,"items":[` + expectedTvShowJson + `]}`

	tvShowResponse = &TvShowResponse{
		Payload: tvShow1,
		Status:  200,
	}

	expectedTvShowResponseJson = `{"data":` + expectedTvShowJson + `}`

	tvShowListResponse = &TvShowListResponse{
		Payload: []*TvShow{tvShow1},
		Status:  200,
	}

	expectedTvShowListResponseJson = `{"data":[` + expectedTvShowJson + `]}`
)

func TestTvShowMarshalling(t *testing.T) {
	data, err := json.Marshal(tvShow1)
	if err != nil {
		t.Fatalf("Error marshalling TvShow: %v", err)
	}

	if string(data) != expectedTvShowJson {
		t.Errorf(testStatusErrorMessage, expectedTvShowJson, string(data))
	}
}

func TestTvShowListMarshalling(t *testing.T) {
	data, err := json.Marshal(tvShowList)
	if err != nil {
		t.Fatalf("Error marshalling TvShowList: %v", err)
	}

	if string(data) != expectedTvShowListJson {
		t.Errorf(testStatusErrorMessage, expectedTvShowListJson, string(data))
	}
}

func TestTvShowResponseMarshalling(t *testing.T) {
	data, err := json.Marshal(tvShowResponse)
	if err != nil {
		t.Fatalf("Error marshalling TvShowResponse: %v", err)
	}

	if string(data) != expectedTvShowResponseJson {
		t.Errorf(testStatusErrorMessage, expectedTvShowResponseJson, string(data))
	}
}

func TestTvShowListResponseMarshalling(t *testing.T) {
	data, err := json.Marshal(tvShowListResponse)
	if err != nil {
		t.Fatalf("Error marshalling TvShowListResponse: %v", err)
	}

	if string(data) != expectedTvShowListResponseJson {
		t.Errorf(testStatusErrorMessage, expectedTvShowListResponseJson, string(data))
	}
}
