package admin

import (
	"net/url"
	"strings"
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
	mediaLibrary := Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Name: "Media Library"})
	mediaLibrary.UseTheme("grid")
	mediaLibrary.UseTheme("media_library")
	mediaLibrary.IndexAttrs("File", "Title")
	Admin.AddResource(db.I18n, &admin.Config{Menu: []string{"Settings"}})
	initWidgets(Admin)

	page := Admin.AddResource(&models.PageSetting{}, &admin.Config{Menu: []string{"Settings"}, Singleton: true})

	page.Meta(&admin.Meta{Name: "TopBackGroundMusic", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
	}})

	page.Meta(&admin.Meta{Name: "TopBackGroundImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 1400, Height: 400},
		},
	}})

	page.Meta(&admin.Meta{Name: "OurBackGroundImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 300, Height: 300},
		},
	}})

	page.Meta(&admin.Meta{Name: "TimeBackGroundImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 1400, Height: 600},
		},
	}})

	page.Meta(&admin.Meta{Name: "StoryBackGroundImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 1400, Height: 1000},
		},
	}})

	page.Meta(&admin.Meta{Name: "UpdateAt",
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			source := resource.(*models.PageSetting)
			source.UpdateAt = time.Now().Format("20060203150405")
		}})

	ourStory := page.Meta(&admin.Meta{Name: "StorySections"}).Resource
	ourStory.Meta(&admin.Meta{Name: "OurStoryImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"pc": {Width: 300, Height: 300},
		},
	}})

	ourStory.Meta(&admin.Meta{Name: "UpdateAt",
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			source := resource.(*models.StorySection)
			source.UpdateAt = time.Now().Format("20060203150405")
		}})

	Admin.RegisterFuncMap("filter_microsite_setting", FilterMicrositeSetting)
}

func FilterMicrositeSetting(requestUrl string, name *admin.Section) bool {
	newRequestUrl, err := url.Parse(requestUrl)
	if err != nil {
		return true
	}

	if microsite_setting_id := newRequestUrl.Query()["microsite_setting_id"]; len(microsite_setting_id) > 0 {
		var sectionId string
		if len(name.Rows) > 0 {
			sectionId = name.Rows[0][0]
			if sectionId == "ID" || sectionId == "Kind" || sectionId == "SerializableMeta" {
				return false
			}
		}

		new_microsite_setting_id := microsite_setting_id[0]
		if strings.HasPrefix(sectionId, new_microsite_setting_id) {
			return false
		}

		microsite_setting_group := newRequestUrl.Query()["microsite_setting_group"]
		if len(microsite_setting_group) > 0 {
			if strings.HasPrefix(sectionId, microsite_setting_group[0]) {
				return false
			}
		}

		return true
	}

	return false
}
