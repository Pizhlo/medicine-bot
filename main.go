package main

import (
	"flag"

	bot "github.com/Pizhlo/medicine-bot/cmd"
)

func main() {
	fileName := flag.String("filename", ".env", "name of config file")
	path := flag.String("path", ".", "path to config file")

	flag.Parse()

	bot.Start(*fileName, *path)
}
