all: build

build: vendor/src
	gb build all

vendor/src:
	gb vendor update --all
