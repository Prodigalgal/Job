package middleware

import (
	"Job/src/beans"
	"Job/src/token"
	"Job/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/storyicon/grbac"
)

func QueryRoles(su *beans.SecurityUser) ([]string, error) {
	// 获取角色
	return []string{"*"}, nil
}

func loadRules() grbac.Rules {
	var rules = grbac.Rules{
		{
			ID: 0,
			Resource: &grbac.Resource{
				Host:   "*",
				Path:   "**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
				ForbiddenRoles:  []string{"*"},
				AllowAnyone:     true,
			},
		},
		{
			ID: 1,
			Resource: &grbac.Resource{
				Host:   "*",
				Path:   "**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
				ForbiddenRoles:  []string{"*"},
				AllowAnyone:     true,
			},
		},
	}
	return rules
}

func Authentication() iris.Handler {
	rules := loadRules()

	rbac, err := grbac.New(grbac.WithRules(rules))
	if err != nil {
		panic(err)
	}

	return func(ctx iris.Context) {
		su := token.GetAccessToken(ctx)
		roles, err := QueryRoles(su)

		if err != nil {
			r := beans.NewRBuilder().
				SetMessage("RBAC ERROR").
				SetCode(iris.StatusInternalServerError).
				SetError(err).
				Build()
			utils.JsonBackR(ctx, r)
			ctx.StopExecution()
			return
		}

		state, err := rbac.IsRequestGranted(ctx.Request(), roles)

		if err != nil {
			r := beans.NewRBuilder().
				SetMessage("RBAC Verify ERROR").
				SetCode(iris.StatusInternalServerError).
				SetError(err).
				Build()
			utils.JsonBackR(ctx, r)
			ctx.StopExecution()
			return
		}

		if !state.IsGranted() {
			r := beans.NewRBuilder().
				SetMessage("RBAC Authentication ERROR").
				SetCode(iris.StatusInternalServerError).
				SetError(err).
				Build()
			utils.JsonBackR(ctx, r)
			ctx.StopExecution()
			return
		}
	}
}
