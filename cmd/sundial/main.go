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
	MAX_NOTES     = 10
)

var (
	// App
	ROOT       string
	SFFont     rl.Font
	appState   = 1
	mousePoint = rl.NewVector2(0.0, 0.0)
	// Background
	background      ui.Background
	backgroundMusic ui.Music
	// Buttons
	addBtn  *ui.Button
	nextBtn *ui.Button
	backBtn *ui.Button
	// Canvas
	notes     []*ui.Note
	drawQueue []*ui.Note
	occupied  = 0
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
	rl.SetTargetFPS(60)
	rl.SetTextLineSpacing(14)

	background = *ui.NewBackground(ui.DynamicTheme(ROOT+"assets/images/backgrounds/"), rl.NewVector2(SCREEN_WIDTH, SCREEN_HEIGHT))

	// Buttons
	nextBtn = ui.NewButton(
		ROOT+"assets/components/button/arrow_right_btn.png",
		ROOT+"assets/components/button/btnsound.wav",
		rl.NewVector2(SCREEN_WIDTH-20, SCREEN_HEIGHT-20), 1,
		func() {
			appState = 2
			fmt.Println("Entering canvas mode")
		},
	)

	addBtn = ui.NewButton(
		ROOT+"assets/components/button/add_btn.png",
		ROOT+"assets/components/button/btnsound.wav",
		rl.NewVector2(SCREEN_WIDTH-20, SCREEN_HEIGHT-20), 1,
		func() {
			newNote := ui.NewNote(
				rl.NewVector2(
					float32(rl.GetRandomValue(300, SCREEN_WIDTH-300)),
					float32(rl.GetRandomValue(100, SCREEN_HEIGHT-200)),
				),
				rl.NewVector2(SCREEN_WIDTH, SCREEN_HEIGHT),
			)
			if occupied < MAX_NOTES {
				notes = append(notes, newNote)
				occupied++
			}
		},
	)

	backBtn = ui.NewButton(
		ROOT+"assets/components/button/arrow_left_btn.png",
		ROOT+"assets/components/button/btnsound.wav",
		rl.NewVector2(0+50, SCREEN_HEIGHT-20), 1,
		func() {
			appState = 1
			fmt.Println("Entering title screen")
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
		nextBtn.Update(mousePoint)
	case 2:
		addBtn.Update(mousePoint)
		backBtn.Update(mousePoint)

		if rl.IsKeyPressed(rl.KeyE) {
			appState = 3
		}
		for _, note := range notes {
			// Expand notes
			if !note.IsExpanded && rl.CheckCollisionPointRec(mousePoint, note.Src) {
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					note.IsExpanded = !note.IsExpanded
					break
				}
			}
			// Close notes
			if note.IsExpanded && !rl.CheckCollisionPointRec(mousePoint, note.Dest) {
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					note.IsExpanded = !note.IsExpanded
					break
				}
			}
		}
	case 3:
		addBtn.Update(mousePoint)
		backBtn.Update(mousePoint)

		if rl.IsKeyPressed(rl.KeyE) {
			appState = 2
		}

		for _, note := range notes {
			// Drag notes
			if !note.IsExpanded && rl.CheckCollisionPointRec(mousePoint, note.Src) {
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
					note.Src.X = mousePoint.X - 0.5*note.Src.Width
					note.Src.Y = mousePoint.Y - 0.5*note.Src.Height
				}
			}
			// Close notes
			if note.IsExpanded && !rl.CheckCollisionPointRec(mousePoint, note.Dest) {
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					note.IsExpanded = !note.IsExpanded
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

		nextBtn.Draw()
	case 2, 3:

		// Canvas

		addBtn.Draw()
		backBtn.Draw()
		for _, note := range notes {
			if note.IsExpanded {
				drawQueue = append(drawQueue, note)
			} else {
				drawQueue = append([]*ui.Note{note}, drawQueue...)
			}
		}
		for _, note := range drawQueue {
			if note.IsExpanded {
				rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.NewColor(0, 0, 0, 150))
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
	nextBtn.Unload()
	backBtn.Unload()
	addBtn.Unload()

	// Unload Audio
	backgroundMusic.Unload()
	rl.CloseAudioDevice()

	// Close the f*ckin app
	rl.CloseWindow()
}
