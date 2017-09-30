package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jinzhu/configor"
	"github.com/theplant/appkit/log"
)

var (
	Logger = log.Default()
	Root   string
)

var Config = struct {
	Port string
	Env  string
	DB   struct {
		Name     string `default:"ilove"`
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		Adapter  string `default:"mysql"`
		User     string `default:"root"`
		Password string `secret:"true"`
	}
}{}

func init() {
	Root = path.Join(os.Getenv("GOPATH"), "/src/github.com/chenxin0723/ilove/")

	os.Setenv("CONFIGOR_ENV_PREFIX", "ILOVE")

	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Printf("Config: ------------%s\n", Config)
	return
}

func IsDraft() bool {
	return strings.HasSuffix(Config.Env, "draft") || Config.Env == "development"
}

func IsProd() bool {
	return Config.Env == "prod"
}
