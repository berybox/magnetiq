package main

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands, &commandInitConfig)
}

var (
	flagFilename = &cli.PathFlag{Name: "filename", Aliases: []string{"f"}, Usage: "Name of output file", Value: "magnetiq.json"}

	commandInitConfig = cli.Command{
		Name:                   "init-config",
		Usage:                  "Creates empty configuration file to be filled by user",
		UseShortOptionHandling: false,
		Flags: []cli.Flag{
			flagFilename,
		},
		Action: func(ctx *cli.Context) error {
			s := settings{
				QBitURL:         "put-qBittorrent-WebUI-URL-here",
				QBitUsername:    "put-qBittorrent-WebUI-Username-here",
				QBitPassword:    "put-qBittorrent-WebUI-Password-here",
				MTLSCertificate: "put-mTLS-p12-certificate-filename-here",
				MTLSPassword:    "put-mTLS-p12-certificate-password-here",
			}

			jsonBytes, err := json.MarshalIndent(s, "", "\t")
			if err != nil {
				return err
			}

			err = os.WriteFile(ctx.Path(flagFilename.Name), jsonBytes, os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
