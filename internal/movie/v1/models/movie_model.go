package models

type MovieList struct {
	ItemCount int32    `json:"item_count"`
	PageCount int16    `json:"page_count"`
	Items     []*Movie `json:"items"`
}

type Movie struct {
	id               string
	MediaId          string   `json:"media_id"`
	Title            string   `json:"title"`
	OriginalTitle    string   `json:"original_title"`
	GenreIds         []int16  `json:"genre_ids"`
	Popularity       float32  `json:"popularity"`
	OriginCountry    []string `json:"origin_country"`
	VoteCount        int32    `json:"vote_count"`
	ReleaseDate      string   `json:"release_date"`
	BackdropPath     string   `json:"backdrop_path"`
	OriginalLanguage string   `json:"original_language"`
	VoteAverage      float32  `json:"vote_average"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
}

type MovieResponse struct {
	Payload      *Movie  `json:"data,omitempty"`
	Status       int     `json:"-"`
	ErrorMessage *string `json:"-"`
}

type MovieListResponse struct {
	Payload      []*Movie `json:"data"`
	Status       int      `json:"-"`
	ErrorMessage *string  `json:"-"`
}
