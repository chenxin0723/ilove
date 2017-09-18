package controllers

import (
	"html/template"

	"github.com/chenxin0723/ilove/config/bindatafs"
	"github.com/chenxin0723/ilove/utils"
	"github.com/gin-gonic/gin"
	"github.com/qor/render"
)

var Render *render.Render

func init() {
	Render = render.New(&render.Config{ViewPaths: []string{"app/views"}})
	Render.SetAssetFS(bindatafs.AssetFS.NameSpace("views"))
}

func SetContext(ctx *gin.Context) {
	funcMap := utils.FuncMap(ctx.Writer, ctx.Request)
	ctx.Set("FuncMap", funcMap)
}

func Home(ctx *gin.Context) {
	FuncMap := ctx.MustGet("FuncMap").(template.FuncMap)
	viewCtx := map[string]interface{}{}
	viewCtx["PageType"] = "Homepage"
	Render.Layout("application").Funcs(FuncMap).Execute("index", viewCtx, ctx.Request, ctx.Writer)
	return
}
