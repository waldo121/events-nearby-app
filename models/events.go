package models

type Event struct {
	Name string
	Description string
	Where Location
	When Duration
	Type string
	Cost string
	Registration string
}

