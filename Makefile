.PHONY: build

build:
	go build -ldflags=-checklinkname=0 -o omni-balance ./cmd
