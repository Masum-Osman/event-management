package dto

import "event_management/models"

type WorkshopsDetails struct {
	Id                int64  `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	StartAt           string `json:"start_at"`
	EndAt             string `json:"end_at"`
	TotalReservations int    `json:"total_reservations"`
}

type WorkshopListResponse struct {
	Id        int64                                  `json:"id"`
	Title     string                                 `json:"title"`
	StartAt   string                                 `json:"start_at"`
	EndAt     string                                 `json:"end_at"`
	Workshops []models.WorkshopDetailsWithoutEventId `json:"workshops"`
}
