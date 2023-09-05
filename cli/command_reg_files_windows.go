package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands, &commandRegFiles)
}

const (
	regFilename   = "magnetiq-register.reg"
	unregFilename = "magnetiq-unregister.reg"
	regFile       = `Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\magnet]
"URL Protocol"=""
@="URL:Magnet link torrent Protocol"
"DefaultIcon"="\"%s\""

[HKEY_CLASSES_ROOT\magnet\shell]

[HKEY_CLASSES_ROOT\magnet\shell\open]

[HKEY_CLASSES_ROOT\magnet\shell\open\command]
@="\"%s\" \"add\"%s \"-l\" \"%%1\""`

	unregFile = `Windows Registry Editor Version 5.00

[-HKEY_CLASSES_ROOT\magnet]`
)

var (
	flagDefaultCategory = &cli.StringFlag{Name: "category", Aliases: []string{"c", "a"}, Usage: "Default category of the newly added torrents"}

	commandRegFiles = cli.Command{
		Name:  "reg-files",
		Usage: "Creates Windows .reg files ('" + regFilename + "' and '" + unregFilename + "') for simple magnet:// protocol (un)registration",
		Flags: []cli.Flag{
			flagDefaultCategory,
		},
		UseShortOptionHandling: false,

		Action: func(ctx *cli.Context) error {
			// Register file
			currPath, err := os.Executable()
			if err != nil {
				return err
			}

			icoPath := strings.TrimSuffix(currPath, filepath.Ext(currPath)) + ".ico"

			currPath = strings.ReplaceAll(currPath, `\`, `\\`)
			icoPath = strings.ReplaceAll(icoPath, `\`, `\\`)
			category := ""

			if ctx.IsSet(flagDefaultCategory.Name) {
				category = " \\\"-a\\\" \\\"" + ctx.String(flagDefaultCategory.Name) + "\\\""
			}

			regFileBytes := fmt.Sprintf(regFile, icoPath, currPath, category)

			err = os.WriteFile(regFilename, []byte(regFileBytes), os.ModePerm)
			if err != nil {
				return err
			}

			// Unregister file
			err = os.WriteFile(unregFilename, []byte(unregFile), os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
