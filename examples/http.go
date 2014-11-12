package main

import (
	spotify "github.com/AbhiAgarwal/go-spotify"
	martini "github.com/go-martini/martini"
)

func main() {
	// Gets command and initializes martini
	commands := spotify.Commands()
	m := martini.Classic()

	// Gets the list of commands to what they do, and appends them
	currentString := ""
	for k := range commands {
		currentString += k + " - " + commands[k] + "\n"
	}

	// Runs the martini server with our given string
	m.Get("/", func() string {
		return currentString
	})

	m.Get("/:name", func(params martini.Params) string {
		if val, ok := commands[params["name"]]; ok {
			// Handle Edge Cases
			if params["name"] == "playPlaylist" || params["name"] == "playTrack" {
				return "Needs a second parameter!"
			}
			// Generate Case will just work
			spotify.Execute(val)
			return "Command has been executed"
		}
		return "Command does not exist"
	})

	m.Get("/:name/:second", func(params martini.Params) string {
		if val, ok := commands[params["name"]]; ok {
			// Handle Edge Cases
			if params["name"] != "playPlaylist" && params["name"] != "playTrack" {
				return "Can't have a second parameter!"
			}
			// Generate Case will just work
			spotify.Execute(spotify.Format(val, params["second"]))
			return "Command has been executed"
		}
		return "Command does not exist"
	})

	m.Run()
}
