all:
	go build src/spotify.go
	sudo mv ./spotify /usr/bin/
