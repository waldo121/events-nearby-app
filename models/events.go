package models

type Event struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Address string `json:"titre_adresse"`
	Latitude string `json:"lat"`
	Longitude string `json:"long"`
	EventStart string `json:"date_debut"`
	EventEnd string `json:"date_fin"`
	Type string `json:"type_evenement"`
	Cost string `json:"cout"`
	Registration string `json:"inscription"`
}

