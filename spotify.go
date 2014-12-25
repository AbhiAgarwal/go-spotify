package spotify

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

// Conversion from io.Reader to string
func ReaderToString(out io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))
	return s
}

func GetValue(keyCommand string) string {
	fullCommand := "tell Application \"Spotify\"" + keyCommand
	c := exec.Command("/usr/bin/osascript", "-e", fullCommand)
	defer c.Wait()
	out, _ := c.StdoutPipe()
	if err := c.Start(); err != nil {
		fmt.Println(keyCommand, "not available")
	}
	return ReaderToString(out)
}

func Format(command, key string) string {
	return fmt.Sprintf(command, key)
}

// Volume needs a seperate command because it needs a conversion from
// io.Reader to string. Also we need the out, _ command so I'll seperate it
// for now.
func GetVolume() string {
	return GetValue("to sound volume as integer")
}

func GetCurrentTrack() string {
	return GetValue("to name of current track")
}

func GetCurrentAlbum() string {
	return GetValue("to album of current track")
}

func GetCurrentArtist() string {
	return GetValue("to artist of current track")
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

func Commands() map[string]string {
	commands := make(map[string]string)
	commands["play"] = "to play"
	commands["playpause"] = "to playpause"
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
	return commands
}
