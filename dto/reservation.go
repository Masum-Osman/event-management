package dto

import "event_management/models"

type ReservationReqBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReservationWithoutWorkshopId struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReservationResponseBody struct {
	Reservation ReservationWithoutWorkshopId         `json:"reservation"`
	Event       models.Events                        `json:"event"`
	Workshop    models.WorkshopDetailsWithoutEventId `json:"workshop"`
}
