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
	SCREEN_WIDTH    = 1920
	SCREEN_HEIGHT   = 1080
	APP_NAME        = "sundial"
	MAX_NOTES       = 10
	MAX_INPUT_CHARS = 100
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
	notes        [MAX_NOTES]*ui.Note
	occupied     = 0
	isEditMode   = false
	letterCounts [MAX_INPUT_CHARS]int
	mode         = map[int]string{
		1: "Title screen",
		2: "Normal mode",
		3: "Edit mode",
		4: "Write mode",
	}
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
				for i := 0; i < MAX_NOTES; i++ {
					if notes[i] == nil {
						notes[i] = newNote
						break
					}
				}
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
	backgroundMusic = *ui.NewMusic(ROOT + "assets/music/North.wav")
	backgroundMusic.Play()

	// Get ROOT
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)
	ROOT = wd[:strings.Index(wd, APP_NAME)+1] + APP_NAME + "/"
	os.Chdir(ROOT)
}

func AppUpdate() {
	backgroundMusic.Update()
	// Mute Music / Unmute music
	if rl.IsKeyPressed(rl.KeyM) && !isEditMode {
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

		if rl.IsKeyPressed(rl.KeyE) && !isEditMode {
			appState = 3
		}
		for _, note := range notes {
			if note != nil {
				// Expand notes
				if !note.IsExpanded && rl.CheckCollisionPointRec(mousePoint, note.Src) {
					if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						note.IsExpanded = !note.IsExpanded
						appState = 4
						break
					}
				}
			}
		}
	case 3:
		addBtn.Update(mousePoint)
		backBtn.Update(mousePoint)

		if rl.IsKeyPressed(rl.KeyE) {
			appState = 2
		}

		for i, note := range notes {
			if note != nil && !note.IsExpanded {
				if rl.CheckCollisionPointRec(mousePoint, note.Src) {
					if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
						// Drag notes
						note.Src.X = mousePoint.X - 0.5*note.Src.Width
						note.Src.Y = mousePoint.Y - 0.5*note.Src.Height
					} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
						// Delete notes
						notes[i] = nil
						letterCounts[i] = 0
						occupied--
					}
				}
			}
		}
	case 4:
		for i, note := range notes {
			if note != nil && note.IsExpanded {
				if rl.CheckCollisionPointRec(mousePoint, note.Dest) {
					if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || isEditMode {
						isEditMode = true
						key := rl.GetKeyPressed()
						if key != 0 {
							fmt.Println(key)
							if key >= 32 && key <= 125 && letterCounts[i] < MAX_INPUT_CHARS {
								note.Content[letterCounts[i]] = byte(key)
								note.Content[letterCounts[i]+1] = '\000'
								letterCounts[i]++
							} else if rl.IsKeyPressed(rl.KeyBackspace) {
								letterCounts[i]--
								note.Content[letterCounts[i]] = '\000'
							} else if rl.IsKeyPressed(rl.KeyEnter) {
								isEditMode = false
								fmt.Println("Exiting input mode")
								break
							}
						}
					}
				}
				// Close notes
				if !rl.CheckCollisionPointRec(mousePoint, note.Dest) {
					if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						note.IsExpanded = !note.IsExpanded
						fmt.Println(note.Content)
						isEditMode = false
						fmt.Println("Exiting input mode")
						appState = 2
						break
					}
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

	// rl.DrawText(strconv.Itoa(appState), 20, 20, 50, rl.White)
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
			if note != nil {
				if !note.IsExpanded {
					note.DrawTextureMini()
				}
			}
		}
	case 4:
		for _, note := range notes {
			if note != nil {
				if note.IsExpanded {
					rl.DrawRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.NewColor(0, 0, 0, 150))
					note.DrawTextureEx()
					rl.DrawTextEx(SFFont, string(note.Content[:]), rl.NewVector2(note.Dest.X+50, note.Dest.Y+70), float32(SFFont.BaseSize)/2, 1, rl.Black)
					break
				}
			}
		}
	default:
	}
	textDims := rl.MeasureTextEx(SFFont, mode[appState], float32(SFFont.BaseSize)/2, 1)
	rl.DrawRectangle(
		int32(SCREEN_WIDTH/2-textDims.X/2-15), int32(SCREEN_HEIGHT-textDims.Y-10),
		int32(textDims.X+30), int32(textDims.Y+10), rl.NewColor(17, 21, 25, 255))

	rl.DrawTextEx(
		SFFont,
		mode[appState],
		rl.NewVector2(SCREEN_WIDTH/2-textDims.X/2, SCREEN_HEIGHT-textDims.Y),
		float32(SFFont.BaseSize)/2, 1, rl.White)

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
