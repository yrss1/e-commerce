package models

type Client struct {
	User
	Balance float64 `json:"balance"`
}
