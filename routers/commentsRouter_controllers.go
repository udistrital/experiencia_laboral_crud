package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:DatoAdicionalTipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["experiencia_laboral_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
