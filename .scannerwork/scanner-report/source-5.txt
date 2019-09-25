package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type DatoAdicionalExperienciaLaboral struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	TipoDatoAdicional  int                 `orm:"column(tipo_dato_adicional)"`
	ExperienciaLaboral *ExperienciaLaboral `orm:"column(experiencia_laboral);rel(fk)"`
	Valor              string              `orm:"column(valor)"`
	Activo             bool                `orm:"column(activo)"`
	FechaCreacion      string              `orm:"column(fecha_creacion);null"`
	FechaModificacion  string              `orm:"column(fecha_modificacion);null"`
}

func (t *DatoAdicionalExperienciaLaboral) TableName() string {
	return "dato_adicional_experiencia_laboral"
}

func init() {
	orm.RegisterModel(new(DatoAdicionalExperienciaLaboral))
}

// AddDatoAdicionalExperienciaLaboral insert a new DatoAdicionalExperienciaLaboral into database and returns
// last inserted Id on success.
func AddDatoAdicionalExperienciaLaboral(m *DatoAdicionalExperienciaLaboral) (id int64, err error) {
	m.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDatoAdicionalExperienciaLaboralById retrieves DatoAdicionalExperienciaLaboral by Id. Returns error if
// Id doesn't exist
func GetDatoAdicionalExperienciaLaboralById(id int) (v *DatoAdicionalExperienciaLaboral, err error) {
	o := orm.NewOrm()
	v = &DatoAdicionalExperienciaLaboral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDatoAdicionalExperienciaLaboral retrieves all DatoAdicionalExperienciaLaboral matches certain condition. Returns empty list if
// no records exist
func GetAllDatoAdicionalExperienciaLaboral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DatoAdicionalExperienciaLaboral))
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

	var l []DatoAdicionalExperienciaLaboral
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

// UpdateDatoAdicionalExperienciaLaboral updates DatoAdicionalExperienciaLaboral by Id and returns error if
// the record to be updated doesn't exist
func UpdateDatoAdicionalExperienciaLaboralById(m *DatoAdicionalExperienciaLaboral) (err error) {
	o := orm.NewOrm()
	v := DatoAdicionalExperienciaLaboral{Id: m.Id}
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "TipoDatoAdicional", "ExperienciaLaboral", "Valor", "Activo", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDatoAdicionalExperienciaLaboral deletes DatoAdicionalExperienciaLaboral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDatoAdicionalExperienciaLaboral(id int) (err error) {
	o := orm.NewOrm()
	v := DatoAdicionalExperienciaLaboral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DatoAdicionalExperienciaLaboral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
