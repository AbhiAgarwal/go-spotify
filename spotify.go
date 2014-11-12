package spotify

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

func GetCurrentTrack() {
	keyCommand := "to name of current track"
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
