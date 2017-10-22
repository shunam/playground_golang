package controllers

import (
	"github.com/astaxie/beego"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

func (this *LoginController) Prepare() {

	this.Layout = "layouts/login.tpl"

	this.Data["Stylesheets"] = []string{
		"https://fonts.googleapis.com/css?family=Roboto:regular,bold,italic,thin,light,bolditalic,black,medium&amp;lang=en",
		"https://fonts.googleapis.com/icon?family=Material+Icons",
		"/css/material.cyan-light_blue.min.css",
		"/css/login.css",
	}

	this.Data["Javascripts"] = []string{
		"/js/material.min.js",
	}
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Get", c.Get)
}

func (c *LoginController) Get() {
	c.TplName = "login/index.tpl"
}
