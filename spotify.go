package main

import (
    "fmt"
    "os/exec"
)

func execute(keyCommand string) {
    fullCommand := "tell Application \"Spotify\"" + keyCommand
    c := exec.Command("/usr/bin/osascript", "-e", fullCommand)
    if err := c.Run(); err != nil {
        fmt.Println(keyCommand, "not available")
    }
}

func examples(commands map[string]string){
    /*
        play current song
            execute(commands["play"])

        playTrack, and then append Track Number
            execute(fmt.Sprintf(commands["playTrack"], "2lFTzUnuGaWlWHJQokjRyb"))
    */
    execute(fmt.Sprintf(commands["playTrack"], "2lFTzUnuGaWlWHJQokjRyb"))
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
    commands["quit"] = "to quit"
    examples(commands)
}
