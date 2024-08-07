package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	Texture rl.Texture2D
	Src     rl.Rectangle
	Dest    rl.Rectangle
	State   int
	Sound   rl.Sound
	Onclick func()
}

func NewButton(texturePath, soundPath string, destPos rl.Vector2, scale float32, onClick func()) *Button {
	texture := rl.LoadTexture(texturePath)
	sound := rl.LoadSound(soundPath)

	return &Button{
		Texture: texture,
		Src:     rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
		Dest:    rl.NewRectangle(destPos.X-float32(texture.Width)*scale, destPos.Y-float32(texture.Height)*scale, float32(texture.Width)*scale, float32(texture.Height)*scale),
		Sound:   sound,
		Onclick: onClick,
	}
}

func (b *Button) Update(mousePoint rl.Vector2) {
	if rl.CheckCollisionPointRec(mousePoint, b.Dest) {
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			b.State = 2
		} else {
			b.State = 1
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			rl.PlaySound(b.Sound)
			if b.Onclick != nil {
				b.Onclick()
			}
		}
	} else {
		b.State = 0
	}
}

func (b *Button) Draw() {
	var color rl.Color
	switch b.State {
	case 0:
		color = rl.White
	case 1:
		color = rl.Yellow
	case 2:
		color = rl.Green
	}
	rl.DrawTextureRec(b.Texture, b.Src, rl.NewVector2(b.Dest.X, b.Dest.Y), color)
}

func (b *Button) Unload() {
	rl.UnloadTexture(b.Texture)
	rl.UnloadSound(b.Sound)
}
