package config

import (
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/configor"
	"github.com/theplant/appkit/log"
)

var (
	Logger = log.Default()
	Root   string
)

var Config = struct {
	Port string
	DB   struct {
		Name     string `default:"asics3"`
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		Adapter  string `default:"mysql"`
		User     string `default:"root"`
		Password string `secret:"true"`
	}
}{}

func init() {
	fmt.Println("###############", Config)
	Root = path.Join(os.Getenv("GOPATH"), "/src/github.com/chenxin0723/ilove/")

	os.Setenv("CONFIGOR_ENV_PREFIX", "ILOVE")

	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Println("###############", Config)

	return
}
