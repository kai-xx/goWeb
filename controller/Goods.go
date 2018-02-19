package controller

import (
	"github.com/devfeel/dotweb"
	"goWeb/model"
	"github.com/gohouse/gorose/utils"
)

func GetGoodsList(context dotweb.Context) error {
	res := model.GetGoodsList(context)

	context.WriteJson(utils.SuccessReturn(res))

	return nil
}
