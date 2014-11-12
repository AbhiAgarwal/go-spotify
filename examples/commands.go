package main

import (
	"fmt"
	spotify "github.com/AbhiAgarwal/go-spotify"
	"os"
	"strconv"
)

func main() {
	commands := make(map[string]string)
	commands["play"] = "to play"
	commands["nextTrack"] = "to next track"
	commands["previousTrack"] = "to previous track"
	commands["pause"] = "to pause"
	commands["playPause"] = "to playpause"
	commands["playTrack"] = "to play track \"spotify:track:%s\""
	commands["playPlaylist"] = "to play track \"spotify:user:ni_co:playlist:%s\""
	commands["repeatOn"] = "to set repeating to true"
	commands["repeatOff"] = "to set repeating to false"
	commands["shuffleOn"] = "to set shuffling to true"
	commands["shuffleOff"] = "to set shuffling to false"
	commands["volumeUp"] = "to set sound volume to %s"
	commands["open"] = "to open"
	commands["quit"] = "to quit"

	if len(os.Args) > 1 {
		if os.Args[1] == "play" {
			if len(os.Args) == 2 {
				spotify.Execute(commands["play"])
			} else {
				spotify.Execute(spotify.Format(commands["playTrack"], os.Args[2]))
			}
		} else if os.Args[1] == "track" {
			fmt.Print(spotify.GetCurrentTrack())
		} else if os.Args[1] == "pause" {
			spotify.Execute(commands["pause"])
		} else if os.Args[1] == "playlist" {
			spotify.Execute(spotify.Format(commands["playPlaylist"], os.Args[2]))
		} else if os.Args[1] == "next" {
			spotify.Execute(commands["nextTrack"])
		} else if os.Args[1] == "previous" {
			spotify.Execute(commands["previousTrack"])
		} else if os.Args[1] == "volume" {
			if len(os.Args) == 3 {
				inputValue, err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				outputValue := strconv.Itoa(inputValue)
				spotify.Execute(spotify.Format(commands["volumeUp"], outputValue))
			} else {
				fmt.Print(spotify.GetVolume())
			}
		} else if os.Args[1] == "up" {
			spotify.ChangeVolume(commands, 10)
		} else if os.Args[1] == "down" {
			spotify.ChangeVolume(commands, -10)
		} else if os.Args[1] == "shuffle" {
			if len(os.Args) == 3 {
				if os.Args[2] == "on" {
					spotify.Execute(commands["shuffleOn"])
				} else if os.Args[2] == "off" {
					spotify.Execute(commands["shuffleOff"])
				}
			}
		} else if os.Args[1] == "repeat" {
			if len(os.Args) == 3 {
				if os.Args[2] == "on" {
					spotify.Execute(commands["repeatOn"])
				} else if os.Args[2] == "off" {
					spotify.Execute(commands["repeatOff"])
				}
			}
		} else if os.Args[1] == "search" {
			if len(os.Args) == 2 {
				fmt.Println("search song <song>")
			} else {
				if os.Args[2] == "song" {
					if len(os.Args) == 4 {
						spotify.SearchTrack(os.Args[3])
					} else {
						fmt.Println("Please enter a song!")
					}
				}
			}
		} else if os.Args[1] == "open" {
			spotify.Execute(commands["open"])
		} else if os.Args[1] == "quit" {
			spotify.Execute(commands["quit"])
		} else {
			fmt.Println("Command not found")
		}
	} else {
		fmt.Println("Spotify Options")
		fmt.Println("   play                   = Start playing Spotify")
		fmt.Println("   track		  = Get your current track")
		fmt.Println("   play <uri>             = Start playing specified Spotify URI")
		fmt.Println("   playlist <uri>         = Start playing playlist Spotify URI")
		fmt.Println("   search song <song>     = Search a particular <song>")
		fmt.Println("   pause                  = Pause Spotify")
		fmt.Println("   next                   = Play next song")
		fmt.Println("   previous               = Play previous song")
		fmt.Println("   shuffle <on/off>       = Shuffle on or off?")
		fmt.Println("   repeat <on/off>        = Repeat on or off?")
		fmt.Println("   volume                 = Get volume of Spotify")
		fmt.Println("   volume <amount>        = Set volume by Amount")
		fmt.Println("   up                     = Increase volume by 10%")
		fmt.Println("   down                   = Decrease volume by 10%")
		fmt.Println("   open                   = Open Spotify")
		fmt.Println("   quit                   = Quit Spotify")
	}
}
