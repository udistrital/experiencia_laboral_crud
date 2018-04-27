// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"experiencia_laboral_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_dedicacion",
			beego.NSInclude(
				&controllers.TipoDedicacionController{},
			),
		),

		beego.NSNamespace("/cargo",
			beego.NSInclude(
				&controllers.CargoController{},
			),
		),

		beego.NSNamespace("/soporte_experiencia_laboral",
			beego.NSInclude(
				&controllers.SoporteExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/dato_adicional_experiencia_laboral",
			beego.NSInclude(
				&controllers.DatoAdicionalExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/tipo_vinculacion",
			beego.NSInclude(
				&controllers.TipoVinculacionController{},
			),
		),

		beego.NSNamespace("/experiencia_laboral",
			beego.NSInclude(
				&controllers.ExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/tipo_experiencia_laboral",
			beego.NSInclude(
				&controllers.TipoExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/dato_adicional_tipo_experiencia_laboral",
			beego.NSInclude(
				&controllers.DatoAdicionalTipoExperienciaLaboralController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
