package gopush

type Push struct {
	Message  string `json:"message"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Url      string `json:"url"`
}
