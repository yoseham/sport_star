// file: controllers/user_controller.go

package controllers

import (
	"app/project/datamodels"
	"app/project/datasource"
	"app/project/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct {
	Ctx iris.Context
	Service services.ISportstarService

}
func (c *IndexController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":"球星库",
			"DataList":dataList,
		},
	}
}
func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path:"/",
		}
	}
	data := c.Service.GetByID(id)
	return mvc.View{
		Name:"info.html",
		Data: iris.Map{
			"Title":"球星库",
			"info":data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	dataList := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":"球星库",
			"DataList":dataList,
		},
	}
}
//http://localhost:8080/clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&datamodels.StarInfo{})
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
	}
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}