package models

type Episode struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type DL struct {
	Fan      string    `json:"fan"`
	Episodes []Episode `json:"episodes"`
}
