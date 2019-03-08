PROJECTNAME := templater
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
PLATFORMS := windows linux darwin

build: 
	GO111MODULE=on CGO_ENABLED=0 \
		go build \
		-ldflags="-s -w" \
		-o $(GOBIN)/kubectl_$(PROJECTNAME) main.go

run:
	GO111MODULE=on go run $(RUNNER)

## clean: Clean build files.
clean:
	go clean
	rm -rf $(GOBIN)/*

## format: Format source code.
format:
	gofmt -w -s .

verify:
	go vet $(RUNNER)

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: help verify format clean run build