package models

type Testimonials struct {
	Id         int    `json:"id"`
	Name       string `json:"name" db:"name"`
	Profession string `json:"profession"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
	Image      string `json:"image"`
}
