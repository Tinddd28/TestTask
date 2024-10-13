package models

type Song struct {
	ID    int    `json:"id" db:"id"`
	Group string `json:"group" db:"group_name"`
	Name  string `json:"name" db:"song_name"`
	Year  string `json:"releaseDate" db:"release_date"`
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

type Verse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type InsertSongDb struct {
	Group string `json:"group" db:"group_name"`
	Name  string `json:"name" db:"song_name"`
	Year  string `json:"releaseDate" db:"release_date"`
	Link  string `json:"link" db:"link"`
}

type SongDetail struct {
	Link string   `json:"link"`
	Year string   `json:"releaseDate"`
	Text []string `json:"text"`
}
