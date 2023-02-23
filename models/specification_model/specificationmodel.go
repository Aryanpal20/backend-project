package specificationmodel

type Specification struct {
	ID                 int    `json:"id"`
	Length             string `json:"length"`             // movie time
	Original_Language  string `json:"original_language"`  // original language hindi, english etc..
	Year_of_production int    `json:"year_of_production"` // launch year
	Director_Name      string `json:"director_name"`      // director name
	Rating             string `json:"rating"`             // rating in starts
	Genres             string `json:"genres"`             //type of movie like action, comedy
	Cast               string `json:"cast"`               // Name of actors
	Movie_ID           int    `json:"movieid"`
}
