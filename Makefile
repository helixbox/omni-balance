.PHONY: build run

build:
	go build -ldflags=-checklinkname=0 -o omni-balance ./cmd

run: build
	./omni-balance -c config_local.yaml -p

