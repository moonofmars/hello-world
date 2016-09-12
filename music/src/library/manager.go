// manager
package library

import "fmt"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Typp   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)} //之前大括号写成了小括号，变成类型转换，报错
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) List() {
	fmt.Println("music detail length:", len(m.musics))
	for i, na := range m.musics {
		fmt.Println("i:", i)
		fmt.Println("music detail:", na)
	}
}
