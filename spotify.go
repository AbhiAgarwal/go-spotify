package main

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"
)

func Execute(keyCommand string) {
	fullCommand := "tell Application \"Spotify\"" + keyCommand
	c := exec.Command("/usr/bin/osascript", "-e", fullCommand)
	defer c.Wait()
	if err := c.Start(); err != nil {
		fmt.Println(keyCommand, "not available")
	}
}

// Volume needs a seperate command because it needs a conversion from
// io.Reader to string. Also we need the out, _ command so I'll seperate it
// for now.
func GetVolume() string {
	keyCommand := "to sound volume as integer"
	fullCommand := "tell Application \"Spotify\"" + keyCommand
	c := exec.Command("/usr/bin/osascript", "-e", fullCommand)
	defer c.Wait()
	out, _ := c.StdoutPipe()
	if err := c.Start(); err != nil {
		fmt.Println(keyCommand, "not available")
	}

	// Conversion from io.Reader to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))
	return s
}

func ChangeVolume(commands map[string]string, volumeAmount int) {
	words := strings.Fields(GetVolume())
	volume := words[0]
	inputValue, err := strconv.Atoi(volume)
	if err != nil {
		fmt.Println(err)
	}
	outputValue := strconv.Itoa(inputValue + volumeAmount)
	Execute(Format(commands["volumeUp"], outputValue))
}

func SearchTrack(trackName string) {
	url := "http://ws.spotify.com/search/1/track.json?q=" + trackName
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	js, err := simplejson.NewJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	songChoices := make([]string, 5)
	for i := 0; i < 5; i++ {
		allRange := js.Get("tracks").GetIndex(i)
		trackName := allRange.Get("name").MustString()
		artistName := allRange.Get("artists").GetIndex(0).Get("name").MustString()
		hrefNumber := allRange.Get("href").MustString()

		fmt.Printf("%d. ", i)
		fmt.Println("Track Name: "+trackName+",", "Artist Name: "+artistName)
		songChoices[i] = hrefNumber
	}

	fmt.Print("Input song choice: ")
	var choiceNumber int
	fmt.Scanf("%d", &choiceNumber)

	if choiceNumber >= 5 || choiceNumber < 0 {
		choiceNumber = 0
	}

	songName := Format("to play track \"%s\"", songChoices[choiceNumber])
	Execute(songName)
}

func Format(command, key string) string {
	return fmt.Sprintf(command, key)
}

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
				Execute(commands["play"])
			} else {
				Execute(Format(commands["playTrack"], os.Args[2]))
			}
		} else if os.Args[1] == "pause" {
			Execute(commands["pause"])
		} else if os.Args[1] == "playlist" {
			Execute(Format(commands["playPlaylist"], os.Args[2]))
		} else if os.Args[1] == "next" {
			Execute(commands["nextTrack"])
		} else if os.Args[1] == "previous" {
			Execute(commands["previousTrack"])
		} else if os.Args[1] == "volume" {
			if len(os.Args) == 3 {
				inputValue, err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				outputValue := strconv.Itoa(inputValue)
				Execute(Format(commands["volumeUp"], outputValue))
			} else {
				fmt.Print(GetVolume())
			}
		} else if os.Args[1] == "up" {
			ChangeVolume(commands, 10)
		} else if os.Args[1] == "down" {
			ChangeVolume(commands, -10)
		} else if os.Args[1] == "shuffle" {
			if len(os.Args) == 3 {
				if os.Args[2] == "on" {
					Execute(commands["shuffleOn"])
				} else if os.Args[2] == "off" {
					Execute(commands["shuffleOff"])
				}
			}
		} else if os.Args[1] == "repeat" {
			if len(os.Args) == 3 {
				if os.Args[2] == "on" {
					Execute(commands["repeatOn"])
				} else if os.Args[2] == "off" {
					Execute(commands["repeatOff"])
				}
			}
		} else if os.Args[1] == "search" {
			if len(os.Args) == 2 {
				fmt.Println("search song <song>")
			} else {
				if os.Args[2] == "song" {
					if len(os.Args) == 4 {
						SearchTrack(os.Args[3])
					} else {
						fmt.Println("Please enter a song!")
					}
				}
			}
		} else if os.Args[1] == "open" {
			Execute(commands["open"])
		} else if os.Args[1] == "quit" {
			Execute(commands["quit"])
		} else {
			fmt.Println("Command not found")
		}
	} else {
		fmt.Println("Spotify Options")
		fmt.Println("   play                   = Start playing Spotify")
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