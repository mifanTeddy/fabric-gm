/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package comm

import (
	"context"
	"crypto/rand"
	"encoding/pem"
	"math/big"

	"github.com/hyperledger/fabric-gm/common/util"
	"github.com/tjfoc/gmsm/sm2"
	tls "github.com/tjfoc/gmtls"
	credentials "github.com/tjfoc/gmtls/gmcredentials"
	"google.golang.org/grpc/peer"
)

// GenerateCertificatesOrPanic generates a a random pair of public and private keys
// and return TLS certificate
func GenerateCertificatesOrPanic() tls.Certificate {
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKey, err := sm2.GenerateKey()
	if err != nil {
		panic(err)
	}
	sn, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		panic(err)
	}
	template := sm2.Certificate{
		KeyUsage:     sm2.KeyUsageKeyEncipherment | sm2.KeyUsageDigitalSignature,
		SerialNumber: sn,
		ExtKeyUsage:  []sm2.ExtKeyUsage{sm2.ExtKeyUsageServerAuth},
	}
	rawBytes, err := sm2.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}
	//privBytes, err := x509.MarshalECPrivateKey(privateKey)
	privBytes, err := sm2.MarshalSm2UnecryptedPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	encodedCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: rawBytes})
	encodedKey := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})
	cert, err := tls.X509KeyPair(encodedCert, encodedKey)
	if err != nil {
		panic(err)
	}
	if len(cert.Certificate) == 0 {
		panic("Certificate chain is empty")
	}
	return cert
}

func certHashFromRawCert(rawCert []byte) []byte {
	if len(rawCert) == 0 {
		return nil
	}
	return util.ComputeSHA256(rawCert)
}

// ExtractCertificateHash extracts the hash of the certificate from the stream
func extractCertificateHashFromContext(ctx context.Context) []byte {
	pr, extracted := peer.FromContext(ctx)
	if !extracted {
		return nil
	}

	authInfo := pr.AuthInfo
	if authInfo == nil {
		return nil
	}

	tlsInfo, isTLSConn := authInfo.(credentials.TLSInfo)
	if !isTLSConn {
		return nil
	}
	certs := tlsInfo.State.PeerCertificates
	if len(certs) == 0 {
		return nil
	}
	raw := certs[0].Raw
	return certHashFromRawCert(raw)
}
