package api

type Error struct {
	StatusCode   int    `json:"status_code"`
	ReasonPhrase string `json:"reason_phrase"`
	Errors       string `json:"errors,omitempty"`
}
