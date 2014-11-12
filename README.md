Go Spotify
==========

Command-line Spotify client and library for Mac. It uses `osascript` to command, and run the commands. Will be building more ontop of this in the future. Just a quick prototype.

Quick and easy to export it as a binary, and to export it to `/usr/bin/`. This binary file consists of an example application that allows you to use the command line `spotify` client that is built ontop of this library. Run:

Use the `make` command, or:

```
go build examples/commands.go
mv ./commands ./spotify
sudo mv ./spotify /usr/bin/
```

Then you can use `spotify` in the command line. This is not done yet - I'm going to be adding things like uri search, etc. 

**Spotify Options**

```   
play                   = Start playing Spotify
track                  = Gets current track
play <uri>             = Start playing specified Spotify URI
playlist <uri>         = Start playing playlist Spotify URI
search song <song>     = Search a particular <song>
pause                  = Pause Spotify
next                   = Play next song
previous               = Play previous song
shuffle <on/off>       = Shuffle on or off?
repeat <on/off>        = Repeat on or off?
volume                 = Get volume of Spotify
volume <amount>        = Set volume by Amount
up                     = Increase volume by 10%
down                   = Decrease volume by 10%
open                   = Open Spotify
quit                   = Quit Spotify
```
