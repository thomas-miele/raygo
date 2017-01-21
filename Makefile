#
# Raygo
#

all: build

build:
	docker run --rm -v ${PWD}:/usr/src/raygo -w /usr/src/raygo golang:latest go build -v

clean:
	rm raygo

perms:
	sudo chown ${USER}:${USER} ./raygo
