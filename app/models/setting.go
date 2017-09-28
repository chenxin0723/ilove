package models

type PageSetting struct {
	TopBackGroundMusic MediaBox
	TopBackGroundImage MediaBox
	OurImage           MediaBox
	UpdateAt           string
}

func (this PageSetting) TopBackGroundMusicUrl() string {
	return this.TopBackGroundMusic.URL("") + "?" + this.UpdateAt
}

func (this PageSetting) TopBackGroundImageUrl() string {
	return this.TopBackGroundImage.URL("pc") + "?" + this.UpdateAt
}
