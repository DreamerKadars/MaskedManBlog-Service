package service

import iris "github.com/kataras/iris/v12"

func Handler(ctx iris.Context) {
	ctx.WriteString("Test Index Handler")
}