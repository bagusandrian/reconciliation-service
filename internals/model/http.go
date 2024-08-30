package model

type GetDummyRequest struct {
	Text       string `json:"text"`
	Wait       int    `json:"wait"`
	StatusCode int    `json:"status_code"`
}
type GetDummyResponse struct {
	Text           string `json:"text"`
	ProcessingTIme string `json:"processing_time"`
}
