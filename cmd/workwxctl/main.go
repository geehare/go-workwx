package main

import (
	"os"

	"go-workwx/cmd/workwxctl/commands"
)

func main() {
	app := commands.InitApp()
	// ignore errors
	_ = app.Run(os.Args)
}
