package config

import (
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

	return
}
