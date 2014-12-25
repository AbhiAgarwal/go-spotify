package main

import (
	spotify "github.com/AbhiAgarwal/go-spotify"
	"github.com/andlabs/ui"
	"log"
)

func main() {
	go ui.Do(gui)
	err := ui.Go()
	if err != nil {
		log.Print(err)
	}
}

func gui() {
	commands := spotify.Commands()
	leftButton := ui.NewButton("<<")
	playButton := ui.NewButton("||")
	rightButton := ui.NewButton(">>")

	leftButton.OnClicked(func() {
		spotify.Execute(commands["previousTrack"])
	})

	playButton.OnClicked(func() {
		spotify.Execute(commands["playpause"])
	})

	rightButton.OnClicked(func() {
		spotify.Execute(commands["nextTrack"])
	})

	stack := ui.NewHorizontalStack(
		leftButton,
		playButton,
		rightButton)

	w := ui.NewWindow("Spotify Client", 100, 25, stack)
	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	w.Show()
}
