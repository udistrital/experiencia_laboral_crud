package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExperienciaLaboral struct {
	Id                     int                     `orm:"column(id);pk"`
	Persona                int                     `orm:"column(persona)"`
	Actividades            string                  `orm:"column(actividades);null"`
	Organizacion           int                     `orm:"column(organizacion)"`
	FechaInicio            time.Time               `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion      time.Time               `orm:"column(fecha_finalizacion);type(date)"`
	TipoDedicacion         *TipoDedicacion         `orm:"column(tipo_dedicacion);rel(fk)"`
	Cargo                  *Cargo                  `orm:"column(cargo);rel(fk)"`
	TipoVinculacion        *TipoVinculacion        `orm:"column(tipo_vinculacion);rel(fk)"`
	TipoExperienciaLaboral *TipoExperienciaLaboral `orm:"column(tipo_experiencia_laboral);rel(fk)"`
}

func (t *ExperienciaLaboral) TableName() string {
	return "experiencia_laboral"
}

func init() {
	orm.RegisterModel(new(ExperienciaLaboral))
}

// AddExperienciaLaboral insert a new ExperienciaLaboral into database and returns
// last inserted Id on success.
func AddExperienciaLaboral(m *ExperienciaLaboral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExperienciaLaboralById retrieves ExperienciaLaboral by Id. Returns error if
// Id doesn't exist
func GetExperienciaLaboralById(id int) (v *ExperienciaLaboral, err error) {
	o := orm.NewOrm()
	v = &ExperienciaLaboral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExperienciaLaboral retrieves all ExperienciaLaboral matches certain condition. Returns empty list if
// no records exist
func GetAllExperienciaLaboral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExperienciaLaboral))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []ExperienciaLaboral
	qs = qs.OrderBy(sortFields...)
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

// UpdateExperienciaLaboral updates ExperienciaLaboral by Id and returns error if
// the record to be updated doesn't exist
func UpdateExperienciaLaboralById(m *ExperienciaLaboral) (err error) {
	o := orm.NewOrm()
	v := ExperienciaLaboral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExperienciaLaboral deletes ExperienciaLaboral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExperienciaLaboral(id int) (err error) {
	o := orm.NewOrm()
	v := ExperienciaLaboral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExperienciaLaboral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
