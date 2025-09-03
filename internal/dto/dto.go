package dto

import "time"

type Request struct {
	URL string `json:"url" binding:"required"`
}

type Response struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
