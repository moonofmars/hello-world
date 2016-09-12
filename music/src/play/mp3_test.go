// mp3_test
package mp

import (
	"fmt"
	"testing"
)

func TestMP3(t *testing.T) {
	fmt.Println("start")
	mp3 := new(MP3Player)
	mp3.Playm("c:\\1.mp3")
}
