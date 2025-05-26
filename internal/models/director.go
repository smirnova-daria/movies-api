package models

type Director struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	BirthYear int    `json:"birth_year"`
	WikiLink  string `json:"wiki_link"`
}
