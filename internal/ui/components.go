package ui

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func AddClock(font rl.Font, SCREEN_WIDTH, SCREEN_HEIGHT int, color rl.Color) {
	dt := time.Now()
	day := fmt.Sprintf("%v, %d tháng %d", dt.Weekday().String(), dt.Day(), dt.Month())
	dayWidth := rl.MeasureTextEx(font, day, float32(font.BaseSize/3), 14)
	clock := fmt.Sprintf("%02d:%02d", dt.Hour(), dt.Minute())
	clockWidth := rl.MeasureTextEx(font, clock, float32(font.BaseSize), 14)
	rl.DrawTextEx(
		font,
		day,
		rl.NewVector2(float32(SCREEN_WIDTH)/2-float32(dayWidth.X)/2, float32(SCREEN_HEIGHT)/2-200),
		float32(font.BaseSize/3),
		14,
		color)

	rl.DrawTextEx(
		font,
		clock,
		rl.NewVector2(float32(SCREEN_WIDTH)/2-float32(clockWidth.X)/2, float32(SCREEN_HEIGHT)/2-150),
		float32(font.BaseSize),
		14,
		color)
}
