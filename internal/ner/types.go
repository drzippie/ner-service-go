package ner

type Entity struct {
	Tag   string `json:"tag"`
	Score string `json:"score"`
	Label string `json:"label"`
}

type ExtractRequest struct {
	Text string `json:"text"`
}

type ExtractResponse struct {
	Entities []Entity `json:"entities"`
}
