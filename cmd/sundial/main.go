package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	ui "github.com/zeann3th/sundial/internal/ui"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
)

var (
	background      rl.Texture2D
	backgroundDest  rl.Rectangle
	backgroundSrc   rl.Rectangle
	SFFont          rl.Font
	appState        = 1
	backgroundMusic rl.Music
	isMusicMuted    bool
)

func main() {
	defer rl.CloseWindow()

	AppStartUp()

	for !rl.WindowShouldClose() {
		AppUpdate()
		AppRender()
	}

	AppShutDown()
}

func init() {
	// Display
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "SunDial")
	rl.SetExitKey(0)
	if rl.IsWindowFocused() {
		rl.SetTargetFPS(60)
	} else {
		rl.SetTargetFPS(15)
	}
	rl.SetTextLineSpacing(14)

	// Audio
	rl.InitAudioDevice()
}

func AppStartUp() {
	// Display
	background = rl.LoadTexture(ui.DynamicTheme())
	backgroundSrc = rl.NewRectangle(0, 0, float32(background.Width), float32(background.Height))
	backgroundDest = rl.NewRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT)
	// Font
	SFFont = rl.LoadFontEx("../../assets/fonts/SF-Pro.ttf", 75, nil, 250)
	// Music
	backgroundMusic = rl.LoadMusicStream("../../assets/music/Ender Lilies - North.wav")
	rl.PlayMusicStream(backgroundMusic)
}

func AppUpdate() {
	rl.UpdateMusicStream(backgroundMusic)
	if rl.IsKeyPressed(rl.KeyM) {
		isMusicMuted = !isMusicMuted
		if isMusicMuted {
			rl.StopMusicStream(backgroundMusic)
		} else {
			rl.PlayMusicStream(backgroundMusic)
		}
	}
}

func AppRender() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	// Title screen
	if appState == 1 {
		rl.DrawTexturePro(background, backgroundSrc, backgroundDest, rl.NewVector2(0, 0), 0, rl.White)
		rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.Color{R: 0, G: 0, B: 0, A: 150})
		ui.AddClock(SFFont, SCREEN_WIDTH, SCREEN_HEIGHT, rl.White)
	}
	switch appState {
	case 1:
		rl.DrawTexturePro(background, backgroundSrc, backgroundDest, rl.NewVector2(0, 0), 0, rl.White)
		rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.Color{R: 0, G: 0, B: 0, A: 150})
		ui.AddClock(SFFont, SCREEN_WIDTH, SCREEN_HEIGHT, rl.White)
	case 2:
	default:
	}
	rl.EndDrawing()
}

func AppShutDown() {
	// Unload texture
	rl.UnloadTexture(background)

	// Unload audio
	rl.UnloadMusicStream(backgroundMusic)
	rl.CloseAudioDevice()
	// Close the f*ckin app
	rl.CloseWindow()
}
