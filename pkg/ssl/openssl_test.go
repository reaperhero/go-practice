package ssl

import (
	"github.com/spacemonkeygo/openssl"
	"math/big"
	"testing"
	"time"
)

func TestOpenssl(t *testing.T) {
	key, err := openssl.GenerateRSAKey(768)
	if err != nil {
		t.Fatal(err)
	}
	info := &openssl.CertificateInfo{
		Serial:       big.NewInt(int64(1)),
		Issued:       0,
		Expires:      24 * time.Hour,
		Country:      "US",
		Organization: "Test",
		CommonName:   "localhost",
	}
	cert, err := openssl.NewCertificate(info, key)
	if err != nil {
		t.Fatal(err)
	}
	if err := cert.Sign(key, openssl.EVP_SHA256); err != nil {
		t.Fatal(err)
	}
}
