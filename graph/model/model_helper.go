package model

type CreatorDecoded struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Arts []*Art `json:"arts"`
}

type ArtDecoded struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      TypeOf `json:"type"`
	CreatorId string `json:"creator"`
}
