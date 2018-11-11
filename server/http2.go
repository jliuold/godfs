package server

import (
	"crypto/tls"
	. "github.com/godfs-io/godfs"
	"github.com/godfs-io/godfs/logger"
	"golang.org/x/net/http2"
	"net/http"
	"strconv"
)

// server 相关变量
var (
	http2Mux = http.NewServeMux()
	http2Url string

	keyFile  string
	certFile string

	serverTlsConfig *tls.Config
)

// 初始化server启动所需的变量
func initServerVars() {
	http2Url = CONFIG.Godfs.Http2.Host + ":" + strconv.Itoa(CONFIG.Godfs.Http2.Port)
	keyFile = CONFIG.Godfs.Http2.KeyPath
	certFile = CONFIG.Godfs.Http2.CertPath
	serverTlsConfig = &tls.Config{
		ClientAuth: tls.RequireAnyClientCert,
	}
}

func StartHttp2() {
	srv := &http.Server{
		Addr:      http2Url,
		Handler:   http2Mux,
		TLSConfig: serverTlsConfig,
	}
	http2.VerboseLogs = false
	http2.ConfigureServer(srv, &http2.Server{
	})
	go func() {
		srv.ListenAndServeTLS(certFile, keyFile)
		logger.Debug("http2 server started on {}", http2Url)
	}()
}
