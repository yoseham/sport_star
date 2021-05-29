package controllers

import (
	"app/project/datamodels"
	"app/project/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)


type AdminController struct {
	Ctx iris.Context
	Service services.ISportstarService
}

func (c *AdminController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":"管理后台",
			"DataList":dataList,
		},
		Layout: "admin/layout.html",
	}
}
func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *datamodels.StarInfo
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
		data = &datamodels.StarInfo{Id:0,}
	} else {
		data = c.Service.GetByID(id)
	}
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title":"管理后台",
			"info":data,
		},
		Layout: "admin/layout.html",
	}
}
func (c *AdminController) PostSave() mvc.Result {
	info := datamodels.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar", "birthday", "height", "weight", "club", "jersy", "country", "moreinfo"})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}
func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
	}
	c.Service.DeleteByID(id)
	return mvc.Response{
		Path: "/admin/",
	}
}