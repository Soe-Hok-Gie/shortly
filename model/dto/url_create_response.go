package dto

import (
	"time"
)

type CreateURLResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type URLResponse struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	LongURL   string    `json:"long_url"`
	HitCount  int64     `json:"hit_count"`
	CreatedAt time.Time `json:"created_at"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type URLListResponse struct {
	Data []URLResponse `json:"data"`
	Meta Pagination    `json:"meta"`
}
