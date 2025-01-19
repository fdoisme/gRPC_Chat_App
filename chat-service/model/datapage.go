package model

type Pagination struct {
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalItems int       `json:"total_items"`
	TotalPages int       `json:"total_pages"`
	RoomId     int       `json:"room_id"`
	RoomName   string    `json:"room_name"`
	Messages   []Message `json:"data"`
}
