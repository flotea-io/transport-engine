package routers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/context/param"
)

func init() {

  beego.GlobalControllerRouter["flt/controllers:SearchController"] = append(beego.GlobalControllerRouter["flt/controllers:SearchController"],
    beego.ControllerComments{
      Method:           "GetAll",
      Router:           `/`,
      AllowHTTPMethods: []string{"get"},
      MethodParams:     param.Make(),
      Filters:          nil,
      Params:           nil})

  beego.GlobalControllerRouter["flt/controllers:SearchController"] = append(beego.GlobalControllerRouter["flt/controllers:SearchController"],
    beego.ControllerComments{
      Method:           "Get",
      Router:           `/:Id`,
      AllowHTTPMethods: []string{"get"},
      MethodParams:     param.Make(),
      Filters:          nil,
      Params:           nil})

}
