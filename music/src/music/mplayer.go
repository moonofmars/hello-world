// mplayer
package main

import (
	"bufio"
	"fmt"
	"library"
	"os"
	"play"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 0
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "find":
		lib.Find(tokens[2])
	case "list":
		lib.List()
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2],
				tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("Usage: lib add <name><artist>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	mp.Play(e.Source, e.Typp)
}

func main() {
	fmt.Println(` Enter following commands to control the player: 
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib 
	play <name> -- Play the specified music `)

	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i <= 100; i++ {
		fmt.Print("Enter command->")
		rawLine, _, _ := r.ReadLine() //unc (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		switch tokens[0] {
		case "lib":
			handleLibCommands(tokens)
		case "play":
			handlePlayCommand(tokens)
		default:
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}

}
