package serializer

type MovieSerializer struct {
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Poster   string `json:"poster"`
}

type MovieResponseSerializer struct {
	Movie MovieSerializer `json:"movie"`
}
