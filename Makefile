all:
	go build examples/commands.go
	mv ./commands ./spotify
	sudo mv ./spotify /usr/bin/
