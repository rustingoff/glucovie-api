package models

type EventModel struct {
	Title  string `json:"title"`
	From   string `json:"from"`
	To     string `json:"to"`
	UserID string `json:"user_id"`
}

type EventResponse struct {
	Title string `json:"title"`
	From  string `json:"from"`
	To    string `json:"to"`
}
