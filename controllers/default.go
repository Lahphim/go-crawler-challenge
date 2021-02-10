package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Layout = "layouts/application.html"
	c.TplName = "main/index.html"
}
