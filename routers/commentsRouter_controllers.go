package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:CargoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:DatoAdicionalExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:RelacionCargosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:SoporteExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoDedicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/experiencia_laboral_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
