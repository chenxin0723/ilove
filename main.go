package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/chenxin0723/ilove/config"
	"github.com/chenxin0723/ilove/config/bindatafs"
	_ "github.com/chenxin0723/ilove/config/db"
	"github.com/chenxin0723/ilove/config/routes"
	"github.com/theplant/appkit/server"
)

func main() {
	var handler http.Handler

	httpLogger := config.Logger.With("origin", "http")
	httpLogger.Debug().Log("msg", "Initializing AIGLE CMS http.Handler instance...")

	handler, _ = routes.Mux(httpLogger)

	middleware := server.Compose(
		server.DefaultMiddleware(httpLogger),
	)

	var compileTemplate = flag.Bool("compile-templates", false, "Compile Templates")
	flag.Parse()
	if *compileTemplate {
		bindatafs.AssetFS.Compile()
		return
	}

	port := fmt.Sprintf(":%s", config.Config.Port)
	fmt.Println("listen on", port)

	err := http.ListenAndServe(port, middleware(handler))
	if err != nil {
		panic(err)
	}
}
