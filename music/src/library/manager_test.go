// manager_test
package library

import (
	"fmt"
	"testing"
)

func TestLib(t *testing.T) {
	fmt.Println("Hello World!")
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed, nil.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}
	m0 := &MusicEntry{
		"1", "My heart will go on", "Celion Dion", "http://qbox.me/24501234", "MP3", //如果最后一个逗号不加：missing ',' before newline in composite literal
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}
}
