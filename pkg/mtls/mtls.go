package mtls

import (
	"crypto/tls"
	"encoding/pem"
	"io"

	"golang.org/x/crypto/pkcs12"
)

// CertFromP12 Creates tls.Certificate usable in http.Client from p12 certificate
func CertFromP12(file io.ReadCloser, password string) (tls.Certificate, error) {
	var ret tls.Certificate

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return ret, err
	}
	defer file.Close()

	pemBlocks, err := pkcs12.ToPEM(fileBytes, password)
	if err != nil {
		return ret, err
	}

	var pemData []byte
	for _, b := range pemBlocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	ret, err = tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
