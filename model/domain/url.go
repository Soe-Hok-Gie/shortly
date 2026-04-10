package domain

import "time"

type URL struct {
	Id       int64
	Code     string
	LongURL  string
	HitCount int64
	CreateAt time.Time
}

type FindURLParams struct {
	UserID int64
	Limit  int
	Offset int
}
