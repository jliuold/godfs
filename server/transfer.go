package server

import (
	. "github.com/godfs-io/godfs"
	"github.com/godfs-io/godfs/logger"
	"github.com/godfs-io/godfs/util"
	"github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	// 文件传输默认的参数名
	file          = "file"
	fileUrl       = "/" + file
	fileSubfix    = "." + file
	tmp           = "tmp"
	fileTmpSubfix = "." + tmp
)

func transfer(formFile multipart.File) {

	// 存储文件
	destDir := filepath.Join(CONFIG.Godfs.File.Storage.Path, util.NowDateDir())
	os.MkdirAll(destDir, os.ModePerm)
	tmpName := uuid.Must(uuid.NewV4()).String() + fileTmpSubfix
	tmpPath := filepath.Join(destDir, tmpName)
	tmpFile, _ := os.Create(tmpPath)
	io.Copy(tmpFile, formFile)
	defer tmpFile.Close()

	// 重命名文件
	tFile, _ := os.Open(tmpPath)
	defer tFile.Close()
	hex := util.Md5Hex(tFile)
	destName := hex + fileSubfix
	destFile := filepath.Join(destDir, destName)
	os.Rename(tmpPath, destFile)
}

func httpUpload(w http.ResponseWriter, r *http.Request) {

	// 解析表单数据
	formFile, _, err := r.FormFile(file)
	if err != nil {
		logger.Error("process http file failed {}", err)
		responseHttp(w, ErrorParam)
		return
	}
	defer formFile.Close()
	transfer(formFile)
	responseHttp(w, OK)
	return
}

func httpIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><form method='post' action='/file'  ENCTYPE='multipart/form-data'><input type='file' name='file'/><input type='submit'/></form></body>"))
}

func init() {

}
