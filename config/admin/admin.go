package admin

import (
	"time"

	"github.com/chenxin0723/ilove/app/models"
	"github.com/chenxin0723/ilove/config/db"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
)

var (
	ActionBar      *action_bar.ActionBar
	ImagesResource *admin.Resource
	Admin          = admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
)

//agileadmin
func init() {
	Admin.SetSiteName("Ilove")
	ActionBar = action_bar.New(Admin)
	ImagesResource = Admin.AddResource(&models.MediaLibrary{}, &admin.Config{Priority: 6})
	ImagesResource.UseTheme("grid")
	ImagesResource.UseTheme("media_library")
	ImagesResource.IndexAttrs("File", "Title")
	ImagesResource.EditAttrs("-TranscriptionFile", "-Localization")
	ImagesResource.NewAttrs("-TranscriptionFile", "-Localization")
	ImagesResource.Meta(&admin.Meta{Name: "TranscriptionFile", Label: "Video Transcript for Accessibility"})
	Admin.AddResource(db.I18n, &admin.Config{Menu: []string{"Settings"}})
	initWidgets(Admin)

	page := Admin.AddResource(&models.PageSetting{}, &admin.Config{Menu: []string{"Settings"}, Singleton: true})

	page.Meta(&admin.Meta{Name: "TopBackGroundImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 1400, Height: 400},
		},
	}})

	page.Meta(&admin.Meta{Name: "UpdateAt",
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			source := resource.(*models.PageSetting)
			source.UpdateAt = time.Now().Format("20060203150405")
		}})

}
