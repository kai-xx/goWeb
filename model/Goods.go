package model

import (
	"github.com/devfeel/dotweb"
	"goWeb/bootstrap"
)

func GetGoodsList (context dotweb.Context) interface{} {
	res, err := bootstrap.DB.Table("sc_gd_prd_goods").Limit(10).Get()

	if err != nil {
		return ""
	}
	return res
}
