package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	ui "github.com/zeann3th/sundial/internal/ui"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
	APP_NAME      = "sundial"
)

var (
	ROOT string
	// Scope: App
	background      rl.Texture2D
	backgroundDest  rl.Rectangle
	backgroundSrc   rl.Rectangle
	SFFont          rl.Font
	appState        = 2
	backgroundMusic rl.Music
	isMusicMuted    bool
	mousePoint      = rl.NewVector2(0.0, 0.0)
	// Scope: Canvas
	NUM_FRAMES = 1
	btn        rl.Texture2D
	btnSrc     rl.Rectangle
	btnDest    rl.Rectangle
	btnState   = 0
	btnSound   rl.Sound
)

func main() {
	defer rl.CloseWindow()

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

	background = rl.LoadTexture(ui.LoadDynamicTheme(ROOT + "assets/images/backgrounds/"))
	backgroundSrc = rl.NewRectangle(0, 0, float32(background.Width), float32(background.Height))
	backgroundDest = rl.NewRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT)
	btn = rl.LoadTexture(ROOT + "assets/components/button/addbtn.png")
	btnSound = rl.LoadSound(ROOT + "assets/components/button/btnsound.wav")

	// Font
	SFFont = rl.LoadFontEx(ROOT+"assets/fonts/SF-Pro.ttf", 50, nil, 250)

	// Audio
	rl.InitAudioDevice()
	backgroundMusic = rl.LoadMusicStream(ROOT + "assets/music/Ender Lilies - North.wav")
	rl.PlayMusicStream(backgroundMusic)

	// Get ROOT
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)
	ROOT = wd[:strings.Index(wd, APP_NAME)] + APP_NAME + "/"
	os.Chdir(ROOT)
}

func AppUpdate() {
	rl.UpdateMusicStream(backgroundMusic)
	// Mute Music / Unmute music
	if rl.IsKeyPressed(rl.KeyM) {
		isMusicMuted = !isMusicMuted
		if isMusicMuted {
			rl.StopMusicStream(backgroundMusic)
		} else {
			rl.PlayMusicStream(backgroundMusic)
		}
	}

	// Check button state
	mousePoint = rl.GetMousePosition()
	if rl.CheckCollisionPointRec(mousePoint, btnDest) {
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			btnState = 2
		} else {
			btnState = 1
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			rl.PlaySound(btnSound)
		}
	} else {
		btnState = 0
	}
}

func AppRender() {
	rl.BeginDrawing()

	// App background
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexturePro(background, backgroundSrc, backgroundDest, rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.Color{R: 0, G: 0, B: 0, A: 150})

	// App state e.g: title screen, note canvas, timetable...
	switch appState {
	case 1:
		// Title screen
		rl.DrawTexturePro(background, backgroundSrc, backgroundDest, rl.NewVector2(0, 0), 0, rl.White)
		rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.Color{R: 0, G: 0, B: 0, A: 100})
		ui.AddClock(SFFont, SCREEN_WIDTH, SCREEN_HEIGHT, rl.White)
	case 2:

		// Canvas

		frameHeight := float32(btn.Height) / float32(NUM_FRAMES)
		btnSrc = rl.NewRectangle(0, 0, float32(btn.Width), frameHeight)
		btnDest = rl.NewRectangle(SCREEN_WIDTH-btnSrc.Width-10, SCREEN_HEIGHT-btnSrc.Height-10, btnSrc.Width, btnSrc.Height)

		if btnState == 0 {
			rl.DrawTextureRec(btn, btnSrc, rl.NewVector2(btnDest.X, btnDest.Y), rl.White)
		} else if btnState == 1 {
			rl.DrawTextureRec(btn, btnSrc, rl.NewVector2(btnDest.X, btnDest.Y), rl.Red)
		} else {
			rl.DrawTextureRec(btn, btnSrc, rl.NewVector2(btnDest.X, btnDest.Y), rl.Green)
		}
	default:
	}

	rl.EndDrawing()
}

func AppShutDown() {
	// Unload texture
	rl.UnloadTexture(background)
	rl.UnloadTexture(btn)

	// Unload audio
	rl.UnloadMusicStream(backgroundMusic)
	rl.UnloadSound(btnSound)
	rl.CloseAudioDevice()

	// Close the f*ckin app
	rl.CloseWindow()
}
