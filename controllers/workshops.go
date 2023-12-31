package controllers

import (
	"event_management/dto"
	models "event_management/models"
	"log"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

// WorkshopsController operations for Workshops
type WorkshopsController struct {
	beego.Controller
}

// URLMapping ...
func (c *WorkshopsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Workshops
// @Param	body		body 	models.Workshops	true		"body for Workshops content"
// @Success 201 {object} models.Workshops
// @Failure 403 body is empty
// @router / [post]
func (c *WorkshopsController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Workshops by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Workshops
// @Failure 403 :id is empty
// @router /:id [get]
func (c *WorkshopsController) GetOne() {
	param := c.Ctx.Input.Param(":id")
	workshopId, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "parse error",
			"status":  "5000",
		}
		c.ServeJSON()
		return
	}

	details, err := models.GetWorkShopDetails(workshopId)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	var response dto.WorkshopsDetails
	response.Id = details.Id
	response.Title = details.Title
	response.Description = details.Description
	response.StartAt = details.StartAt
	response.EndAt = details.EndAt
	response.TotalReservations = details.TotalReservations

	c.Data["json"] = response
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Workshops
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Workshops
// @Failure 403
// @router / [get]
func (c *WorkshopsController) GetAll() {
	param := c.Ctx.Input.Query("eventId")
	eventId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "parse error",
			"status":  "5000",
		}
		c.ServeJSON()
		return
	}

	event, err := models.GetWorkshopsById(eventId)
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	var workShopListResponse dto.WorkshopListResponse
	workShopListResponse.Id = event.Id
	workShopListResponse.Title = event.Title
	workShopListResponse.StartAt = event.StartAt
	workShopListResponse.EndAt = event.EndAt

	workshopList, err := models.GetWorkShopListByEventId(int(eventId))
	if err != nil {
		log.Println(err)
		c.Data["json"] = map[string]string{
			"message": "sql error",
			"status":  "5001",
		}
		c.ServeJSON()
		return
	}

	workShopListResponse.Workshops = *workshopList

	c.Data["json"] = workShopListResponse
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Workshops
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Workshops	true		"body for Workshops content"
// @Success 200 {object} models.Workshops
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WorkshopsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Workshops
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WorkshopsController) Delete() {

}
