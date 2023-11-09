package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"event_management/queries"

	"github.com/beego/beego/v2/client/orm"
)

type Workshops struct {
	Id          int64
	EventId     int64
	StartAt     string `orm:"size(128)"`
	EndAt       string `orm:"size(128)"`
	Title       string `orm:"size(128)"`
	Description string `orm:"size(128)"`
}

type WorkshopsDetailsResponse struct {
	Workshops
	TotalReservations int `json:"total_reservations"`
}

func init() {
	orm.RegisterModel(new(Workshops))
}

// AddWorkshops insert a new Workshops into database and returns
// last inserted Id on success.
func AddWorkshops(m *Workshops) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWorkshopsById retrieves Workshops by Id. Returns error if
// Id doesn't exist
func GetWorkshopsById(id int64) (v *Workshops, err error) {
	o := orm.NewOrm()
	v = &Workshops{Id: id}
	if err = o.QueryTable(new(Workshops)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWorkshops retrieves all Workshops matches certain condition. Returns empty list if
// no records exist
func GetAllWorkshops(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Workshops))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Workshops
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateWorkshops updates Workshops by Id and returns error if
// the record to be updated doesn't exist
func UpdateWorkshopsById(m *Workshops) (err error) {
	o := orm.NewOrm()
	v := Workshops{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWorkshops deletes Workshops by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWorkshops(id int64) (err error) {
	o := orm.NewOrm()
	v := Workshops{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Workshops{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetWorkShopDetails(workshopId int) (*WorkshopsDetailsResponse, error) {
	o := orm.NewOrm()

	var v WorkshopsDetailsResponse

	err := o.Raw(queries.GetWorkShopDetails, workshopId).QueryRow(&v)
	if err != nil {
		log.Println("Failed to fetch data from mysql:", err)
		return nil, err
	}

	return &v, nil
}
