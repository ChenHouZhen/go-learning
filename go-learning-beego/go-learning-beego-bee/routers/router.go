package routers

import (
	"go-learning/go-learning-beego/go-learning-beego-bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
