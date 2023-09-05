package main

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"

	"github.com/berybox/magnetiq/pkg/magnet"
	"github.com/berybox/magnetiq/pkg/mtls"
	"github.com/berybox/magnetiq/pkg/qbit"
	"github.com/gen2brain/beeep"
	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands, &commandAddTorrent)
}

var (
	errInvalidMagnetLink    = errors.New("invalid magnet link")
	msgTorrentAdded         = "Added successfully"
	magnetlinkQueryFilename = "dn"
)

var (
	flagMagnetLink = &cli.StringFlag{Name: "link", Aliases: []string{"l", "u"}, Usage: "Magnet link of the torrent", Required: true}
	flagCategory   = &cli.StringFlag{Name: "category", Aliases: []string{"a"}, Usage: "Assign category of the torrent"}
	flagConfig     = &cli.PathFlag{Name: "config", Aliases: []string{"c"}, Usage: "qBittorrent WebUI configuration file", Value: "magnetiq.json"}

	commandAddTorrent = cli.Command{
		Name:                   "add-torrent",
		Aliases:                []string{"add"},
		Usage:                  "Adds torrent to the qBittorrent queue",
		UseShortOptionHandling: false,
		Flags: []cli.Flag{
			flagMagnetLink,
			flagCategory,
			flagConfig,
		},
		Action: func(ctx *cli.Context) error {
			sets, err := loadSettings(ctx.String(flagConfig.Name))
			if err != nil {
				return err
			}

			httpClient := &http.Client{}

			if sets.MTLSCertificate != "" {
				certfn, err := toAbs(sets.MTLSCertificate)
				if err != nil {
					return err
				}

				file, err := os.Open(certfn)
				if err != nil {
					return err
				}

				cert, err := mtls.CertFromP12(file, sets.MTLSPassword)
				if err != nil {
					return err
				}

				httpClient.Transport = &http.Transport{
					TLSClientConfig: &tls.Config{
						Certificates: []tls.Certificate{cert},
					},
				}
			}

			q, err := qbit.Login(httpClient, sets.QBitURL, sets.QBitUsername, sets.QBitPassword)
			if err != nil {
				return err
			}

			l, err := magnet.LinkFromString(ctx.String(flagMagnetLink.Name))
			if err != nil {
				return err
			}

			if !l.HasParam(magnetlinkQueryFilename) {
				return errInvalidMagnetLink
			}

			err = q.AddTorrent([]string{l.String()}, ctx.String(flagCategory.Name))
			if err != nil {
				return err
			}

			fn := l.GetParam(magnetlinkQueryFilename)
			imgfn, _ := toAbs(imageFilename)
			err = beeep.Notify(fn, msgTorrentAdded, imgfn)
			if err != nil {
				return err
			}

			return nil
		},
	}
)
