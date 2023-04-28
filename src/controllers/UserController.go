package controllers

import (
	"Job/src/beans"
	"Job/src/services"
	"Job/src/token"
	"Job/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/x/errors"
)

func Register(ctx iris.Context) {
	user := new(beans.User)
	err := ctx.ReadJSON(user)

	if err != nil {
		r := beans.NewRBuilder().
			SetMessage("submit info error").
			SetCode(iris.StatusBadRequest).
			SetError(err).
			Build()
		utils.JsonBackR(ctx, r)
		return
	}

	insert, err := services.UserServiceImpl{}.Register(user)

	if err != nil {
		r := beans.NewRBuilder().
			SetMessage("register error").
			SetCode(iris.StatusInternalServerError).
			SetError(err).
			Build()
		utils.JsonBackR(ctx, r)
		return
	}
	if insert == 0 {
		r := beans.NewRBuilder().
			SetMessage("register error").
			SetCode(iris.StatusInternalServerError).
			SetError(errors.New("Insert Error")).
			Build()
		utils.JsonBackR(ctx, r)
		return
	}

	tokenPair, err := token.GenerateTokenPair(user)
	if err != nil {
		r := beans.NewRBuilder().
			SetMessage("token generate error").
			SetCode(iris.StatusInternalServerError).
			SetError(errors.New("token generate error")).
			Build()
		utils.JsonBackR(ctx, r)
		return
	}

	r := beans.NewRBuilder().
		SetMessage("register success").
		SetData(map[string]interface{}{"token": tokenPair}).
		Build()
	utils.JsonBackR(ctx, r)
}
