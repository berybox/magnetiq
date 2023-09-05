package qbit

/*
API Documentation at:
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)
*/

import (
	"net/http"
	"net/url"
	"path"
)

const (
	// UserAgent User-Agent Header used for requests
	UserAgent = "MagnetiQ"
)

var (
	bodyOK = []byte("Ok.")
)

//Client Main structure for qBittorrent WebUI interaction
type Client struct {
	HTTPClient *http.Client
	URL        string
	SID        *http.Cookie
}

//JoinURL join parts of URL for API calls
func JoinURL(base, suffix string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, suffix)

	return u.String(), nil
}
