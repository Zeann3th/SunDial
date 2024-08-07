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
	ROOT            string
	background      ui.Background
	SFFont          rl.Font
	appState        = 1
	backgroundMusic ui.Music
	mousePoint      = rl.NewVector2(0.0, 0.0)
	addBtn          ui.Button
	notes           []*ui.Note
	enterBtn        ui.Button
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

	background = *ui.NewBackground(ui.DynamicTheme(ROOT+"assets/images/backgrounds/"), rl.NewVector2(SCREEN_WIDTH, SCREEN_HEIGHT))

	// Buttons
	addBtn = *ui.NewButton(
		ROOT+"assets/components/button/addbtn.png",
		ROOT+"assets/components/button/btnsound.wav",
		rl.NewVector2(SCREEN_WIDTH-20, SCREEN_HEIGHT-20), 1,
		func() {
			newNote := ui.NewNote(
				rl.NewVector2(
					float32(rl.GetRandomValue(300, SCREEN_WIDTH-300)),
					float32(rl.GetRandomValue(100, SCREEN_HEIGHT-100)),
				),
				rl.NewVector2(SCREEN_WIDTH, SCREEN_HEIGHT),
			)
			notes = append(notes, newNote)
		},
	)

	enterBtn = *ui.NewButton(
		ROOT+"assets/components/button/enterbtn.png",
		ROOT+"assets/components/button/btnsound.wav",
		rl.NewVector2(SCREEN_WIDTH-20, SCREEN_HEIGHT-20), 1,
		func() {
			appState = 2
			fmt.Println("Entering canvas mode")
		},
	)

	// Font
	SFFont = rl.LoadFontEx(ROOT+"assets/fonts/SF-Pro.ttf", 50, nil, 250)

	// Audio
	rl.InitAudioDevice()
	backgroundMusic = *ui.NewMusic(ROOT + "assets/music/Ender Lilies - North.wav")
	backgroundMusic.Play()

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
	backgroundMusic.Update()
	// Mute Music / Unmute music
	if rl.IsKeyPressed(rl.KeyM) {
		backgroundMusic.ToggleMute()
	}

	// Check button state
	mousePoint = rl.GetMousePosition()
	switch appState {
	case 1:
		enterBtn.Update(mousePoint)
	case 2:
		addBtn.Update(mousePoint)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			for _, note := range notes {
				if rl.CheckCollisionPointRec(mousePoint, note.Src) {
					note.IsExpanded = !note.IsExpanded
					break
				}
				if !rl.CheckCollisionPointRec(mousePoint, note.Dest) {
					note.IsExpanded = false
					break
				}
			}
		}

	}
}

func AppRender() {
	rl.BeginDrawing()

	// App background
	rl.ClearBackground(rl.RayWhite)
	background.DrawWithOverlay()

	// App state e.g: title screen, note canvas, timetable...
	switch appState {
	case 1:
		// Title screen
		ui.NewClock(SFFont, SCREEN_WIDTH, SCREEN_HEIGHT, rl.White)

		enterBtn.Draw()
	case 2:

		// Canvas

		addBtn.Draw()
		for _, note := range notes {
			if note.IsExpanded {
				note.DrawTextureEx()
			} else {
				note.DrawTextureMini()
			}
		}
	default:
	}

	rl.EndDrawing()
}

func AppShutDown() {
	// Unload texture
	background.Unload()
	enterBtn.Unload()
	addBtn.Unload()

	// Unload Audio
	backgroundMusic.Unload()
	rl.CloseAudioDevice()

	// Close the f*ckin app
	rl.CloseWindow()
}
