package client

import (
	"crypto/tls"
	"crypto/x509"
	. "github.com/godfs-io/godfs"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

// client相关变量
var (
	certFile string
	keyFile  string

	transport  *http2.Transport
	httpClient http.Client
)

// 初始化client相关的变量
func initClientVars() {
	caCert, _ := ioutil.ReadFile(CONFIG.Godfs.Http2.CertPath)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(nil)
	}

	transport = &http2.Transport{
		AllowHTTP: false, //充许非加密的链接
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			//ClientCAs:          caCertPool,
			Certificates: []tls.Certificate{clientCert},
		},
	}
	httpClient = http.Client{Transport: transport}
}
