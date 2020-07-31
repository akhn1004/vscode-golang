package tasks


type Task struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	Done bool   `json:"done,omitempty"`
}
