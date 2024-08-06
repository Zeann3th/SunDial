package ui

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func AddClock(font rl.Font, SCREEN_WIDTH, SCREEN_HEIGHT int, color rl.Color) {
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

func AddButton() {}
