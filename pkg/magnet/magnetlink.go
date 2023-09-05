package magnet

import (
	"net/url"
)

// Link MagnetLink for storing Magnet URI
type Link struct {
	URL *url.URL
}

// LinkFromString Creates magent Link from URI string
func LinkFromString(magnetlink string) (Link, error) {
	var ret Link
	u, err := url.Parse(magnetlink)
	ret.URL = u
	return ret, err
}

// GetParam gets parameter of magnet link given by name
func (l Link) GetParam(name string) string {
	return l.URL.Query().Get(name)
}

// HasParam checks if magnet link contains parameter
func (l Link) HasParam(name string) bool {
	if l.GetParam(name) == "" {
		return false
	}
	return true
}

func (l Link) String() string {
	return l.URL.String()
}
