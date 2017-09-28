package admin

import (
	"html/template"

	"github.com/chenxin0723/ilove/app/models"
	"github.com/qor/admin"
	"github.com/qor/widget"
)

var Widgets *widget.Widgets

func initWidgets(adm *admin.Admin) {

	if Widgets == nil {
		// PreviewAssets: []string{"/system/static/" + config.GetAsset("container.css"), "/system/static/" + config.GetAsset("vendor_container.js"), "/system/static/" + config.GetAsset("container.js")}}
		Widgets = widget.New(&widget.Config{})
		Widgets.RegisterViewPath("github.com/chenxin0723/ilove/app/views/widgets")

		Widgets.WidgetSettingResource = Admin.AddResource(&models.QorWidgetSetting{}, &admin.Config{Name: "New Container"})
		Widgets.WidgetSettingResource.IndexAttrs("PreviewIcon", "Name", "Description")

		adm.AddResource(Widgets, &admin.Config{Menu: []string{"Containers"}, Priority: 5, Invisible: true})
		Widgets.RegisterFuncMap("Raw", func(val string) template.HTML {
			return template.HTML(val)
		})

		type CommonTest struct {
		}

		header := adm.NewResource(&CommonTest{})
		Widgets.RegisterWidget(&widget.Widget{
			Name:      "CommonTest",
			Templates: []string{"widget"},
			Setting:   header,
			Context: func(context *widget.Context, setting interface{}) *widget.Context {
				return context
			},
		})

	}
}
