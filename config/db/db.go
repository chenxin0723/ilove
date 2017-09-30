package db

import (
	"errors"
	"fmt"

	"github.com/chenxin0723/ilove/app/models"
	"github.com/chenxin0723/ilove/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/media"
	"github.com/qor/media_library"
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
		I18n = i18n.New(database.New(DB))
		migrateDB(DB)
	} else {
		fmt.Println("can't link to db ---", err)
	}
}

func migrateDB(db *gorm.DB) (err error) {
	db.AutoMigrate(&media_library.MediaLibrary{}, &models.PageSetting{}, &models.QorWidgetSetting{}, &models.StorySection{})

	common_widgets := []string{"CommonTest"}
	for _, widget_name := range common_widgets {
		var widget models.QorWidgetSetting
		if db.Where("name = ?", widget_name).Find(&widget).RecordNotFound() {
			widget.Name = widget_name
			widget.WidgetType = widget_name
			widget.GroupName = widget_name
			widget.Kind = widget_name
			widget.Shared = true
			db.Create(&widget)
		}
	}
	return nil
}
