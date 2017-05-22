// @APIVersion 1.0.0
// @Title algorithm-service API
// @Description algorithm-service only serve algorithm
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app-service/algorithm-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/algorithm",
			beego.NSInclude(
				&controllers.AlgorithmController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
