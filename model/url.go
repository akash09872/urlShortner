package model

type URL struct {
	FullUrl string `json:"url"`
}
type Shortened struct {
	Short string `json:"short"`
}
type URLResponse struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
}
