package dto

type EventsDetails struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	StartAt        string `json:"start_at"`
	EndAt          string `json:"end_at"`
	TotalWorkshops int    `json:"total_workshops"`
}
