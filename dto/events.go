package dto

import "event_management/models"

type EventsDetails struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	StartAt        string `json:"start_at"`
	EndAt          string `json:"end_at"`
	TotalWorkshops int    `json:"total_workshops"`
}

type paginate struct {
	Total       int64 `json:"total"`
	PerPage     int64 `json:"per_page"`
	TotalPages  int64 `json:"total_pages"`
	CurrentPage int   `json:"current_page"`
}
type EventListResponse struct {
	Events     []models.Events `json:"events"`
	Pagination paginate        `json:"pagination"`
}
