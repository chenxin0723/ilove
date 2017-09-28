package controllers

import (
	"fmt"
	"html/template"
	"time"

	"github.com/chenxin0723/ilove/app/models"
	"github.com/chenxin0723/ilove/config"
	"github.com/chenxin0723/ilove/config/admin"
	"github.com/chenxin0723/ilove/config/bindatafs"
	"github.com/chenxin0723/ilove/config/db"
	"github.com/chenxin0723/ilove/utils"
	"github.com/gin-gonic/gin"
	"github.com/qor/i18n"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/render"
	"github.com/qor/widget"
)

var Render *render.Render

func init() {
	Render = render.New(&render.Config{ViewPaths: []string{"app/views"}})
	Render.SetAssetFS(bindatafs.AssetFS.NameSpace("views"))
}

func SetContext(ctx *gin.Context) {
	funcMap := utils.FuncMap(ctx.Writer, ctx.Request)
	funcMap["T"] = func(key string, args ...interface{}) template.HTML {
		if len(args) == 0 {
			args = []interface{}{key}
		} else {
			args = append([]interface{}{key}, args...)
		}

		localeI18n := db.I18n.Fallbacks("en-us").(*i18n.I18n)
		var inlineEdit bool
		if config.IsDraft() {
			inlineEdit = true
		}
		return inline_edit.InlineEdit(localeI18n, "en-us", inlineEdit)(key, args...)
	}

	funcMap["get_setting"] = func() interface{} {
		var page_setting models.PageSetting
		db.DB.First(&page_setting)
		return page_setting
	}
	funcMap["inlineEdit"] = func() bool {
		return admin.ActionBar.EditMode(ctx.Writer, ctx.Request)
	}
	options := map[string]interface{}{}
	options["Request"] = ctx.Request
	widgetContext := admin.Widgets.NewContext(&widget.Context{
		DB:         db.DB,
		InlineEdit: admin.ActionBar.EditMode(ctx.Writer, ctx.Request),
		Options:    options,
	})

	funcMap["CommonTest"] = func() template.HTML {
		fmt.Println("#################", widgetContext.Render("CommonTest", "CommonTest"))
		return widgetContext.Render("CommonTest", "CommonTest")
	}

	ctx.Set("FuncMap", funcMap)
}

func Home(ctx *gin.Context) {
	FuncMap := ctx.MustGet("FuncMap").(template.FuncMap)

	viewCtx := map[string]interface{}{}
	if config.IsDraft() {
		qorAdminCssJs := template.HTML(`<script type="text/javascript" src="/system/qor_vendor.js"></script><link rel="stylesheet" href="/system/qor_admin.css"><script defer async type="text/javascript" src="/system/qor_admin.js"></script>`)
		viewCtx["ActionBar"] = qorAdminCssJs + admin.ActionBar.Render(ctx.Writer, ctx.Request)
	}

	viewCtx["PageType"] = "Homepage"
	viewCtx["Now"] = time.Now().Format("Mon Jan 2 . 2006")
	Render.Layout("application").Funcs(FuncMap).Execute("index", viewCtx, ctx.Request, ctx.Writer)
	return
}
