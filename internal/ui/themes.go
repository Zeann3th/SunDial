package ui

import "time"

const (
	prefix    = "../../assets/backgrounds/"
	DAWN      = "dawn.png"
	MORNING   = "morning.png"
	AFTERNOON = "afternoon.png"
	NIGHT     = "night.png"
	SPACE     = "space.png"
)

var currentTheme = prefix + MORNING

func DynamicTheme() string {
	currentTheme = prefix
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
