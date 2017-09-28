package models

import (
	"time"

	"github.com/qor/widget"
)

type QorWidgetSetting struct {
	widget.QorWidgetSetting
	DeletedAt *time.Time
}
