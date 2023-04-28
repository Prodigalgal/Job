package utils

import (
	"Job/src/beans"
	"github.com/kataras/iris/v12"
)

func JsonBackR(ctx iris.Context, r *beans.R) {
	_ = ctx.JSON(r)
}
