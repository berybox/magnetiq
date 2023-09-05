package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type settings struct {
	QBitURL         string `json:"qbit_url"`
	QBitUsername    string `json:"qbit_username"`
	QBitPassword    string `json:"qbit_password"`
	MTLSCertificate string `json:"mtls_certificate"`
	MTLSPassword    string `json:"mtls_password"`
}

func loadSettings(filename string) (settings, error) {
	filename, err := toAbs(filename)
	if err != nil {
		return settings{}, err
	}

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return settings{}, err
	}

	var ret settings
	err = json.Unmarshal(fileBytes, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func toAbs(filename string) (string, error) {
	if !filepath.IsAbs(filename) {
		exe, err := os.Executable()
		if err != nil {
			return "", err
		}
		filename = filepath.Clean(filepath.Join(filepath.Dir(exe), filename))
	}

	return filename, nil
}
