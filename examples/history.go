package main

import (
	"fmt"
	spotify "github.com/AbhiAgarwal/go-spotify"
)

// Still under implementation
func main() {
	currentTrack := spotify.GetCurrentTrack()
	fmt.Print(currentTrack)
}
