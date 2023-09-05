package qbit

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

//API methods
const (
	APIMethodAddTorrent = "/api/v2/torrents/add"
)

//AddTorrent add magnet link torrent to connected client
func (q *Client) AddTorrent(magnetLinks []string, category string) error {
	torrentURL, err := JoinURL(q.URL, APIMethodAddTorrent)
	if err != nil {
		return err
	}

	reqBody, contentType, err := createMultipartBody(magnetLinks, category)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, torrentURL, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Referer", q.URL)
	req.Header.Set("Content-Type", contentType)

	req.AddCookie(q.SID)

	resp, err := q.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if bytes.Compare(body, bodyOK) != 0 {
		return fmt.Errorf("Adding magnet link failed: %s", body)
	}

	return nil
}

// createMultipartBody creates multipart which is required by qBittorrent WebUI. Returns body, content type and error
func createMultipartBody(magnetLinks []string, category string) (io.Reader, string, error) {
	var reqBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&reqBody)
	fieldWriter, err := multipartWriter.CreateFormField("urls")
	if err != nil {
		return nil, "", err
	}

	for _, link := range magnetLinks {
		fmt.Fprintln(fieldWriter, link)
	}

	fieldWriter, err = multipartWriter.CreateFormField("category")
	if err != nil {
		return nil, "", err
	}

	fmt.Fprintln(fieldWriter, category)

	multipartWriter.Close()

	return &reqBody, multipartWriter.FormDataContentType(), nil
}
