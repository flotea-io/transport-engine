/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*
* Framework License
*
* @APIVersion 1.0.0
* @Title beego Test API
* @Description beego has a very cool tools to autogenerate documents for your API
* @Contact astaxie@gmail.com
* @TermsOfServiceUrl http://beego.me/
* @License Apache 2.0
* @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
*/

package routers

import (
	"flt/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/search",
			beego.NSInclude(
				&controllers.SearchController{},
			),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
