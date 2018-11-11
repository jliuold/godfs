package server

import (
	"encoding/json"
	. "github.com/godfs-io/godfs"
	"github.com/godfs-io/godfs/logger"
	"net/http"
)

const (
	// HTTP contentType键
	ContentType = "content-type"

	ApplicationJson = "application/json;charset=utf-8"
)

// server 相关变量
var (
	httpMux = http.NewServeMux()
	httpUrl string
)

// 开启HTTP服务
func StartHttp() {
	go func() {
		http.ListenAndServe(httpUrl, httpMux)
		logger.Debug("http server started on {}", httpUrl)
	}()
}

// 响应Http请求，写入响应体
func responseHttp(w http.ResponseWriter, resp *Resp) {
	bytes, _ := json.Marshal(resp)
	w.Header().Set(ContentType, ApplicationJson)
	w.Write(bytes)
}

func addHandler() {
	httpMux.HandleFunc("/index", httpIndex)
	httpMux.HandleFunc(fileUrl, httpUpload)
}
