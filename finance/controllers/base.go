package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {

	this.Layout = "layouts/dashboard.tpl"

	this.Data["Stylesheets"] = []string{
		"https://fonts.googleapis.com/css?family=Roboto:regular,bold,italic,thin,light,bolditalic,black,medium&amp;lang=en",
		"https://fonts.googleapis.com/icon?family=Material+Icons",
		"/css/material.cyan-light_blue.min.css",
		"/css/styles.css",
	}

	this.Data["Javascripts"] = []string{
		"/js/material.min.js",
	}
}
