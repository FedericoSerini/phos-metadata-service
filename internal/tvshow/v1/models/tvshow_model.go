package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TvShowList struct {
	ItemCount int32     `json:"item_count"`
	PageCount int16     `json:"page_count"`
	Items     []*TvShow `json:"items"`
}

type TvShow struct {
	id                  primitive.ObjectID
	BackdropPath        string               `json:"backdrop_path"`
	CreatedBy           []*CreatedBy         `json:"created_by"`
	EpisodeRuntime      []int32              `json:"episode_run_time"`
	FirstAirDate        string               `json:"first_air_date"`
	Genres              []*Genre             `json:"genres"`
	Homepage            string               `json:"homepage"`
	MediaId             string               `json:"media_id" bson:"media_id"`
	InProduction        bool                 `json:"in_production"`
	Languages           []string             `json:"languages"`
	LastAirDate         string               `json:"last_air_date"`
	LastEpisodeToAir    *LastEpisodeToAir    `json:"last_episode_to_air"`
	Name                string               `json:"name"`
	NextEpisodeToAir    string               `json:"next_episode_to_air"`
	Networks            []*Network           `json:"networks"`
	NumberOfEpisodes    int16                `json:"number_of_episodes"`
	NumberOfSeasons     int16                `json:"number_of_seasons"`
	OriginCountry       []string             `json:"origin_country"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalName        string               `json:"original_name"`
	Overview            string               `json:"overview"`
	Popularity          float32              `json:"popularity"`
	PosterPath          string               `json:"poster_path"`
	ProductionCompanies []*ProductionCompany `json:"production_companies"`
	Seasons             []*Season            `json:"seasons"`
	Episodes            []*Episode           `json:"episodes"`
	Status              string               `json:"status"`
	Type                string               `json:"type"`
	VoteAverage         float32              `json:"vote_average"`
	VoteCount           int32                `json:"vote_count"`
}

type CreatedBy struct {
	Id          int32  `json:"id"`
	CreditId    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int8   `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type Genre struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type LastEpisodeToAir struct {
	Id             int32   `json:"id"`
	EpisodeNumber  int16   `json:"episode_number"`
	AirDate        string  `json:"air_date"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int16   `json:"season_number"`
	ShowId         int32   `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int32   `json:"vote_count"`
}

type Network struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCompany struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type Season struct {
	Id           int32  `json:"id"`
	AirDate      string `json:"air_date"`
	EpisodeCount int16  `json:"episode_count"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int16  `json:"season_number"`
}

type Episode struct {
	MediaId       string `json:"media_id" bson:"media_id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Overview      string `json:"overview"`
	PosterPath    string `json:"poster_path"`
	EpisodeNumber int16  `json:"episode_number"`
}

type TvShowResponse struct {
	Payload      *TvShow `json:"data,omitempty"`
	Status       int     `json:"-"`
	ErrorMessage *string `json:"-"`
}

type TvShowListResponse struct {
	Payload      []*TvShow `json:"data"`
	Status       int       `json:"-"`
	ErrorMessage *string   `json:"-"`
}
