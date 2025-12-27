package dto

type TopLinkResponse struct {
	Code     string `json:"code"`
	LongURL  string `json:"long_url"`
	HitCount int64  `json:"hit_count"`
}
