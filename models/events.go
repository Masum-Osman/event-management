package models

import (
	"errors"
	"event_management/queries"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Events struct {
	Id      int64  `json:"id"`
	Title   string `orm:"size(128)" json:"title"`
	StartAt string `orm:"size(128)" json:"start_at"`
	EndAt   string `orm:"size(128)" json:"end_at"`
}

type EventDetailsResponse struct {
	Workshops
	TotalWorkshops int `json:"total_workshops"`
}

func init() {
	orm.RegisterModel(new(Events))
}

// AddEvents insert a new Events into database and returns
// last inserted Id on success.
func AddEvents(m *Events) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEventsById retrieves Events by Id. Returns error if
// Id doesn't exist
func GetEventsById(id int64) (v *Events, err error) {
	o := orm.NewOrm()
	v = &Events{Id: id}
	if err = o.QueryTable(new(Events)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEvents retrieves all Events matches certain condition. Returns empty list if
// no records exist
func GetAllEvents(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Events))
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

	var l []Events
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

// UpdateEvents updates Events by Id and returns error if
// the record to be updated doesn't exist
func UpdateEventsById(m *Events) (err error) {
	o := orm.NewOrm()
	v := Events{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEvents deletes Events by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEvents(id int64) (err error) {
	o := orm.NewOrm()
	v := Events{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Events{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetEventDetails(eventId int) (*EventDetailsResponse, error) {
	o := orm.NewOrm()

	var v EventDetailsResponse

	err := o.Raw(queries.GetEventDetails, eventId).QueryRow(&v)
	if err != nil {
		log.Println("Failed to fetch data from mysql:", err)
		return nil, err
	}

	return &v, nil
}
