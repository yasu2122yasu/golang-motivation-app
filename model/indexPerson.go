package model

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Saving string `json:"saving"`
}
