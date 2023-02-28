package moviemodel

import detail "movie/models/specification_model"

type Movie struct {
	ID            int                    `json:"id"`
	Movie_Name    string                 `json:"movie_name"`
	Specification []detail.Specification `gorm:"ForeignKey:Movie_ID"`
}
