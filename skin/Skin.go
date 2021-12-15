package skin

import "log"

type Skin struct {
	colors []SkinColor
}

type ColorPair int16

func NewSkin() *Skin {
	s := &Skin{}
	return s
}

func (self *Skin) Register(group string, name string, pairIndex ColorPair) {
	color := SkinColor{
		group:     group,
		name:      name,
		pairIndex: int16(pairIndex),
	}
	self.colors = append(self.colors, color)
}

func (self *Skin) GetColor(group string, name string) ColorPair {
	for _, c := range self.colors {
		if c.group == group && c.name == name {
			return ColorPair(c.pairIndex)
		}
	}
	return 0
}

func (self *Skin) Dump() {
	log.Printf("--- Skin Dump ---")
	for _, c := range self.colors {
		log.Printf("pairIndex = %d group = %s name = %s", c.pairIndex, c.group, c.name)
	}
}
