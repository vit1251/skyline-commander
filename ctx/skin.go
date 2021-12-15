package ctx

import skin "github.com/vit1251/skyline-commander/skin"

var mainSkin *skin.Skin

func GetSkin() *skin.Skin {
	return mainSkin
}

func SetSkin(s *skin.Skin) {
	mainSkin = s
}
