package dto

type TopLinkResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
	HitCount int64  `json:"hit_count"`
}
