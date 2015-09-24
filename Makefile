
all: commands

commands: examples/commands.go
	go build examples/commands.go

install: all
	mv commands /usr/bin/spotify

.PHONY = clean
clean:
	rm commands
