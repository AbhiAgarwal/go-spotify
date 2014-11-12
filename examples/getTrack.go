package main

import (
	"fmt"
	spotify "github.com/AbhiAgarwal/go-spotify"
)

func main() {
	currentTrack := spotify.GetCurrentTrack()
	fmt.Print(currentTrack)
}
