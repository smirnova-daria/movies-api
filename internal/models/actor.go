package models

type Actor struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	BirthYear int    `json:"birth_year"`
	Oscars    int    `json:"oscars"`
	WikiLink  string `json:"wiki_link"`
}
