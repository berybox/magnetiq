package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/urfave/cli/v2"
)

const (
	imageFilename = "magnetiq.png"
)

var (
	commands = []*cli.Command{}
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: false,
		Usage:                  "MagnetiQ - Torrent adding for qBittorrent WebUI",
		Commands:               commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		imgfn, _ := toAbs(imageFilename)
		beeep.Notify("ERROR", fmt.Sprintf("%s", err), imgfn)
		panic(err)
	}
}
