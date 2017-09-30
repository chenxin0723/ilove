package models

import "github.com/jinzhu/gorm"

type StorySection struct {
	gorm.Model
	PageSettingID       uint
	PageSetting         PageSetting
	OurStoryImage       MediaBox
	OurStoryDate        string
	OurStoryDescription string
	UpdateAt            string
}

type PageSetting struct {
	gorm.Model
	TopBackGroundMusic   MediaBox
	TopBackGroundImage   MediaBox
	OurBackGroundImage   MediaBox
	TimeBackGroundImage  MediaBox
	StoryBackGroundImage MediaBox
	StorySections        []StorySection `sql:"type:text;"`
	UpdateAt             string
}

func (this PageSetting) TopBackGroundMusicUrl() string {
	return this.TopBackGroundMusic.URL() + "?" + this.UpdateAt
}

func (this PageSetting) TopBackGroundImageUrl() string {
	return this.TopBackGroundImage.URL("pc") + "?" + this.UpdateAt
}

func (this PageSetting) OurBackGroundImageUrl() string {
	return this.OurBackGroundImage.URL("pc") + "?" + this.UpdateAt
}
func (this PageSetting) TimeBackGroundImageUrl() string {
	return this.TimeBackGroundImage.URL("pc") + "?" + this.UpdateAt
}

func (this PageSetting) StoryBackGroundImageUrl() string {
	return this.StoryBackGroundImage.URL("pc") + "?" + this.UpdateAt
}

func (this StorySection) OurStoryImageUrl() string {
	return this.OurStoryImage.URL("pc") + "?" + this.UpdateAt
}
