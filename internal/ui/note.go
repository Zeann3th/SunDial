package ui

import rl "github.com/gen2brain/raylib-go/raylib"

var palette = map[int]rl.Color{
	0: rl.Red,
	1: rl.Green,
	2: rl.Blue,
	3: rl.Yellow,
}

type Note struct {
	Src        rl.Rectangle
	Dest       rl.Rectangle
	Color      rl.Color
	Content    []string
	IsExpanded bool
	Onclick    func()
}

func NewNote(notePos, screenDims rl.Vector2) *Note {
	baseSize := 200
	scale := 4.0
	return &Note{
		Src:   rl.NewRectangle(notePos.X, notePos.Y, float32(baseSize), float32(baseSize)),
		Dest:  rl.NewRectangle(screenDims.X/2-float32(baseSize/2)*float32(scale), screenDims.Y/2-float32(baseSize/2)*float32(scale), float32(baseSize)*float32(scale), float32(baseSize)*float32(scale)),
		Color: palette[int(rl.GetRandomValue(0, 3))],
	}
}

func (n *Note) DrawTextureMini() {
	rl.DrawRectangle(int32(n.Src.X), int32(n.Src.Y), int32(n.Src.Width), int32(n.Src.Height), n.Color)
}

func (n *Note) DrawTextureEx() {
	rl.DrawRectangle(int32(n.Dest.X), int32(n.Dest.Y), int32(n.Dest.Width), int32(n.Dest.Height), n.Color)
}
