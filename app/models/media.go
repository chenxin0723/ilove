package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/media_library"
)

type MediaLibrary struct {
	gorm.Model
	File              media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/media-libraries/{{primary_key}}/{{column}}.{{extension}}"`
	TranscriptionFile media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/media-libraries/{{primary_key}}/{{column}}.{{extension}}"`
}

type MediaBox struct {
	media_library.MediaBox `sql:"type:text;"`
}

func (mediaBox MediaBox) Description() string {
	if len(mediaBox.Files) > 0 {
		return mediaBox.Files[0].Description
	}
	return ""
}

func (mediaBox MediaBox) SetDescription(des string) {
	if len(mediaBox.Files) > 0 {
		mediaBox.Files[0].Description = des
	}
	return
}
