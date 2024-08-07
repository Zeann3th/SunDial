package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Music struct {
	Stream  rl.Music
	IsMuted bool
}

func NewMusic(filePath string) *Music {
	return &Music{
		Stream:  rl.LoadMusicStream(filePath),
		IsMuted: false,
	}
}

func (m *Music) Play() {
	rl.PlayMusicStream(m.Stream)
}

func (m *Music) Update() {
	rl.UpdateMusicStream(m.Stream)
}

func (m *Music) ToggleMute() {
	m.IsMuted = !m.IsMuted
	if m.IsMuted {
		rl.PauseMusicStream(m.Stream)
	} else {
		rl.ResumeMusicStream(m.Stream)
	}
}

func (m *Music) Restart() {
	rl.StopMusicStream(m.Stream)
	rl.PlayMusicStream(m.Stream)
}

func (m *Music) Unload() {
	rl.UnloadMusicStream(m.Stream)
}
