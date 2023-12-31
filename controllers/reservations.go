package controllers

import (
	"encoding/json"
	"event_management/dto"
	"event_management/models"
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

// ReservationsController operations for Reservations
type ReservationsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReservationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Reservations
// @Param	body		body 	models.Reservations	true		"body for Reservations content"
// @Success 201 {object} models.Reservations
// @Failure 403 body is empty
// @router / [post]
func (c *ReservationsController) Post() {
	var payload dto.ReservationReqBody

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "parse error",
			"status":  "5000",
		}
		c.ServeJSON()
		return
	}

	reservation, err := models.GetReservationsByNameAndEmail(payload.Name, payload.Email)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	var reservationResponse dto.ReservationResponseBody
	reservationResponse.Reservation.Id = int(reservation.Id)
	reservationResponse.Reservation.Name = reservation.Name
	reservationResponse.Reservation.Email = reservation.Email

	workshop, err := models.GetWorkshopsById(reservation.WorkshopId)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	reservationResponse.Workshop.Id = workshop.Id
	reservationResponse.Workshop.Title = workshop.Title
	reservationResponse.Workshop.Description = workshop.Description
	reservationResponse.Workshop.StartAt = workshop.StartAt
	reservationResponse.Workshop.EndAt = workshop.EndAt

	event, err := models.GetEventsById(workshop.EventId)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	reservationResponse.Event = *event

	c.Data["json"] = reservationResponse
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description get Reservations by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Reservations
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ReservationsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Reservations
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Reservations
// @Failure 403
// @router / [get]
func (c *ReservationsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Reservations
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Reservations	true		"body for Reservations content"
// @Success 200 {object} models.Reservations
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ReservationsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Reservations
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ReservationsController) Delete() {

}
