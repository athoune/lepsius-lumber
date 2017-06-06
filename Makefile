export GOPATH:=$(shell pwd)/gopath


build: gopath/src/github.com/athoune/lepsius-lumber vendor
	go build github.com/athoune/lepsius-lumber

gopath/src/github.com/athoune:
	mkdir -p gopath/src/github.com/athoune

gopath/src/github.com/athoune/lepsius-lumber: gopath/src/github.com/athoune
	ln -sf $(shell pwd)  gopath/src/github.com/athoune/lepsius-lumber

vendor:
	glide install

clean:
	rm -rf vendor gopath
