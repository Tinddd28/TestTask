package models

type Song struct {
	ID    int    `json:"id" db:"id"`
	Group string `json:"group" db:"author_name"`
	Name  string `json:"name" db:"song_name"`
	Year  int    `json:"release_date" db:"release_date"`
	Link  string `json:"link" db:"link"`
}

type RequestSong struct {
	Group string `json:"group"`
	Name  string `json:"name"`
}

type ResponseSong struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
