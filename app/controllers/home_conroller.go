package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/chenxin0723/ilove/config/admin"
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
	if isDraft(ctx.Request) {
		qorAdminCssJs := template.HTML(`<script type="text/javascript" src="/system/qor_vendor.js"></script><link rel="stylesheet" href="/system/qor_admin.css"><script defer async type="text/javascript" src="/system/qor_admin.js"></script>`)
		viewCtx["ActionBar"] = qorAdminCssJs + admin.ActionBar.Render(ctx.Writer, ctx.Request)
	}

	viewCtx["PageType"] = "Homepage"
	Render.Layout("application").Funcs(FuncMap).Execute("index", viewCtx, ctx.Request, ctx.Writer)
	return
}

func isDraft(r *http.Request) bool {
	return !strings.Contains(r.Host, "draft-")
}
