package qbit

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//API methods
const (
	APIMethodLogin = "api/v2/auth/login"
)

var (
	loginFailBody = []byte("Fails.")
)

//Login Login to qBittorrent WebUI using custom http client
func Login(client *http.Client, qBitURL, username, password string) (Client, error) {
	var ret Client
	ret.HTTPClient = client
	ret.URL = qBitURL

	loginURL, err := JoinURL(ret.URL, APIMethodLogin)
	if err != nil {
		return ret, err
	}

	loginData := url.Values{}
	loginData.Add("username", username)
	loginData.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, loginURL, strings.NewReader(loginData.Encode()))
	if err != nil {
		return ret, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Referer", ret.URL)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	if bytes.Compare(body, bodyOK) != 0 {
		return ret, fmt.Errorf("Login failed (%s). Invalid credentials?", body)
	}

	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == "SID" {
			ret.SID = c
		}
	}

	if ret.SID == nil {
		return ret, fmt.Errorf("Could not obtain session cookie")
	}

	return ret, nil
}
