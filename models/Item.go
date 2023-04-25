package models

type Item struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         string  `json:"price"`
	Rating        float64 `json:"rating"`
	RatingCounter int     `json:"ratingCounter"`
	Quantity      int     `json:"quantity"`
}
