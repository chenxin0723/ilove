// Package routes is a commen place to put all applicatioin routes.
package routes

import (
	"context"
	"html/template"
	"net/http"
	"strings"

	"github.com/chenxin0723/ilove/app/controllers"
	"github.com/chenxin0723/ilove/config"
	"github.com/chenxin0723/ilove/config/admin"
	"github.com/chenxin0723/ilove/config/db"
	"github.com/gin-gonic/gin"
	"github.com/qor/wildcard_router"
	"github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
	"github.com/theplant/ec/admin_auth"
)

var WildcardRouter = wildcard_router.New()

// Mux initializes the application's HTTP handler.
func Mux(l log.Logger) (http.Handler, error) {

	engine := gin.New()

	engine.RedirectTrailingSlash = false

	mux := http.NewServeMux()
	router := engine.Group("", controllers.SetContext)
	router.GET("/", controllers.Home)

	publicDir := http.Dir(strings.Join([]string{config.Root, "public"}, "/"))
	var fs = http.FileServer(publicDir)
	mux.Handle("/system/", fs)
	mux.Handle("/media-libraries/", fs)
	mux.Handle("/system2/", fs)

	WildcardRouter.MountTo("/", mux)
	admin.Admin.MountTo("/admin", mux)
	WildcardRouter.AddHandler(engine)

	ab := admin_auth.Start(l, db.DB, admin.Admin, []string{})
	ab.AuthLoginOKPath = "/"
	mux.Handle("/auth/", ab.NewRouter())

	middleware := server.Compose(
		// WithWidgetContext(),
		WithRenderActionBar(),
		DraftMiddleWare(),
	)

	return middleware(mux), nil
}
func DraftMiddleWare() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !config.IsDraft() {
				if strings.HasPrefix(r.RequestURI, "/auth/") || strings.HasPrefix(r.RequestURI, "/admin") {
					http.NotFound(w, r)
					return
				}
			}
			h.ServeHTTP(w, r)
		})
	}
}

func WithRenderActionBar() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var actionBar template.HTML
			qorAdminCssJs := template.HTML(`<script type="text/javascript" src="/cms/qor_vendor.js"></script><link rel="stylesheet" href="/cms/qor_admin.css"><script defer async type="text/javascript" src="/cms/qor_admin.js"></script>`)
			actionBar = qorAdminCssJs + admin.ActionBar.Render(w, r)
			newCtx := r.WithContext(context.WithValue(r.Context(), "ActionBar", actionBar))
			h.ServeHTTP(w, newCtx)
		})
	}
}

func InlineEdit(w http.ResponseWriter, req *http.Request, actionBarEditMode bool) (inline_edit bool) {
	if !actionBarEditMode {
		return
	}
	if is_prod_site := req.Header.Get("x-cms-production"); is_prod_site != "" {
		return
	}
	// sinline_edit = actionBarEditMode && config.IsDraft()
	inline_edit = actionBarEditMode
	return
}

type basicAuth struct {
	h http.Handler
}

func (ba *basicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, p, _ := r.BasicAuth()
	if u == "theplant" && p == "theplant works" {
		ba.h.ServeHTTP(w, r)
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}
