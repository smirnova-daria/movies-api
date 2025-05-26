package models

type Movie struct {
	ID            int      `json:"id"`
	TitleOriginal string   `json:"original_title"`
	TitleRussian  string   `json:"russian_title"`
	Year          int      `json:"year"`
	Description   string   `json:"description"`
	Genre         Genre    `json:"genre"`
	Actors        []Actor  `json:"main_actors"`
	Director      Director `json:"director"`
}
