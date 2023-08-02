package gamecore

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontManager interface {
	Default()
}

type fontManager struct {
	control *control
	fonts   map[string]font.Face
}

const (
	defaulltFont    = "default"
	defaultFontSize = 16
	defaultFontDPI  = 72
)

func createFontManager(control *control) *fontManager {
	m := &fontManager{
		control: control,
		fonts:   make(map[string]font.Face),
	}
	m.initDefaultFont()
	return m
}

func (m *fontManager) getDefault() font.Face {
	return m.fonts[defaulltFont]
}

func (m *fontManager) initDefaultFont() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	fontFace, err := opentype.NewFace(
		tt, &opentype.FaceOptions{
			Size:    defaultFontSize,
			DPI:     defaultFontDPI,
			Hinting: font.HintingVertical,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	m.fonts[defaulltFont] = fontFace
}
