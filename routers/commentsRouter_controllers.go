package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["app-service/algorithm-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
