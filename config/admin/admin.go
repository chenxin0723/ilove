package admin

import (
	"github.com/chenxin0723/ilove/config/db"
	"github.com/jinzhu/inflection"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/media/media_library"
	"github.com/qor/publish2"
	"github.com/qor/qor"
)

var (
	ActionBar *action_bar.ActionBar
	Admin     = admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
)

//agileadmin
func init() {
	Admin.SetSiteName("Ilove")
	ActionBar = action_bar.New(Admin)
	mediaLibrary := Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Name: "Media Library"})
	mediaLibrary.UseTheme("grid")
	mediaLibrary.UseTheme("media_library")
	mediaLibrary.IndexAttrs("File", "Title")
	inflection.AddUncountable("Container Library")
}
