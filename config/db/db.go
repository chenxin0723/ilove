package db

import (
	"errors"
	"fmt"

	"github.com/chenxin0723/ilove/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/media"
	"github.com/qor/publish2"
)

var (
	DB               *gorm.DB
	VisibleModeOffDB *gorm.DB
	I18n             *i18n.I18n
)

func init() {
	var err error
	var db *gorm.DB

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?clientFoundRows=true&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	} else {
		panic(errors.New("Database adapter is not supported"))
	}

	if err == nil {
		DB = db
		DB.LogMode(true)
		VisibleModeOffDB = db.Set(publish2.VisibleMode, publish2.ModeOff)
		publish2.RegisterCallbacks(DB)
		media.RegisterCallbacks(DB)
		// migrateDB(DB)
	} else {
		fmt.Println("can't link to db")
	}
}

// func migrateDB(db *gorm.DB) (err error) {
// 	db.AutoMigrate(&models.QorWidgetSetting{}, &media_library.MediaLibrary{}, &banner_editor.QorBannerEditorSetting{}, &models.PageBuilder{})
// 	return nil
// }
