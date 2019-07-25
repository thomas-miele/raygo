#
# Raygo
#

NAME = raygo

all: run

run: build
	./$(NAME) | display

build:
	go build

clean:
	go clean
