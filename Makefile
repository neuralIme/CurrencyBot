.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot main.go bot.go cbrClient.go

run: build
	./.bin/bot