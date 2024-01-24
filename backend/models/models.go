package models

type Artist struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Birthdate        string   `json:"birthdate"`
	ActiveSince      string   `json:"activeSince"`
	Genres           []string `json:"genres"`
	Labels           []string `json:"labels"`
	Website          string   `json:"website"`
	Image            string   `json:"image"`
	ShortDescription string   `json:"shortDescription"`
}
