package main

import (
	"Job/src/controllers"
	"Job/src/middleware"
	"github.com/kataras/iris/v12"
)

func main() {
	App := iris.New()
	App.UseRouter(middleware.CORS())

	userApi := App.Party("/user")
	{
		userApi.Post("/register", controllers.Register)
	}

	err := App.Listen("0.0.0.0:8024", iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		return
	}
}
