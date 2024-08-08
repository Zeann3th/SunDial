package ui

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DAWN      = "dawn.png"
	MORNING   = "morning.png"
	AFTERNOON = "afternoon.png"
	NIGHT     = "night.png"
	SPACE     = "space.png"
)

type Background struct {
	Texture rl.Texture2D
	Src     rl.Rectangle
	Dest    rl.Rectangle
	Music   rl.Music
}

func NewBackground(texturePath string, screenDims rl.Vector2) *Background {
	texture := rl.LoadTexture(texturePath)
	return &Background{
		Texture: texture,
		Src:     rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
		Dest:    rl.NewRectangle(0, 0, screenDims.X, screenDims.Y),
	}
}

func (bg *Background) Draw() {
	rl.DrawTexturePro(bg.Texture, bg.Src, bg.Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func (bg *Background) DrawWithOverlay() {
	bg.Draw()
	rl.DrawRectangle(0, 0, int32(bg.Dest.Width), int32(bg.Dest.Height), rl.NewColor(0, 0, 0, 150))
}

func (bg *Background) Unload() {
	rl.UnloadTexture(bg.Texture)
}

func DynamicTheme(folderPath string) string {
	var currentTheme string
	currentTheme = folderPath
	hours, _, _ := time.Now().Clock()
	if hours >= 0 && hours < 6 {
		currentTheme += DAWN
	} else if hours >= 6 && hours < 12 {
		currentTheme += MORNING
	} else if hours >= 12 && hours < 18 {
		currentTheme += AFTERNOON
	} else {
		currentTheme += NIGHT
	}
	return currentTheme
}

func NewClock(font rl.Font, SCREEN_WIDTH, SCREEN_HEIGHT int, color rl.Color) {
	dt := time.Now()
	// Day. e.g: Tuesday, August 6th
	day := fmt.Sprintf("%v, %s %dth", dt.Weekday().String(), dt.Month().String(), dt.Day())
	dayFontSize := float32(font.BaseSize / 2)
	dayWidth := rl.MeasureTextEx(font, day, dayFontSize, 14)

	rl.DrawTextEx(
		font,
		day,
		rl.NewVector2(float32(SCREEN_WIDTH)/2-float32(dayWidth.X)/2, float32(SCREEN_HEIGHT)/2-200),
		dayFontSize,
		14,
		color)

	// Clock. e.g: 19:09
	clock := fmt.Sprintf("%02d:%02d", dt.Hour(), dt.Minute())
	clockFontSize := float32(font.BaseSize * 3 / 2)
	clockWidth := rl.MeasureTextEx(font, clock, clockFontSize, 14)

	rl.DrawTextEx(
		font,
		clock,
		rl.NewVector2(float32(SCREEN_WIDTH)/2-float32(clockWidth.X)/2, float32(SCREEN_HEIGHT)/2-150),
		clockFontSize,
		14,
		color)
}
