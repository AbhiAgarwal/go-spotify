all:
	go build spotify.go
	sudo mv ./spotify /usr/bin/
