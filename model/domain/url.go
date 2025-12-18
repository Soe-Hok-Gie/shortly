package domain

import "time"

type URL struct {
	Id       int64
	Code     string
	LongURL  string
	HitCount int64
	CreateAt time.Time
}
