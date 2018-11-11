package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/godfs-io/godfs/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//计算文件的Md5值，返回计算结果
func Md5(reader io.Reader) []byte {
	md5hash := md5.New()
	io.Copy(md5hash, reader)
	sum := md5hash.Sum(nil)
	return sum
}

//计算文件的MD5值，返回其16进制
func Md5Hex(reader io.Reader) string {
	return hex.EncodeToString(Md5(reader))
}

func Sha1(reader io.Reader) []byte {
	h := sha1.New()
	io.Copy(h, reader)
	sum := h.Sum(nil)
	return sum;
}

func Sha1Hex(reader io.Reader) string {
	return hex.EncodeToString(Sha1(reader))
}

func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		logger.Error("get current path error {}", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func NowDateDir() string {
	return time.Now().Format("2006/01/02")
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
