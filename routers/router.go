package routers

import (
	"go_code/chapter13/testcase02/lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
