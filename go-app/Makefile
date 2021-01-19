help:
	@echo ""
	@echo "usage: make COMMAND"
	@echo ""
	@echo "Commands:"
	@echo "setup"
	@echo "build"
	@echo "run"
	@echo ""

setup:
	npm install

build:
	npm run sass
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run:
	./gotest


