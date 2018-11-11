package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("/tmp", "godfs_test_1.tmp")
	ioutil.WriteFile(tmpFile.Name(), []byte("abc"), 600)
	println(Md5(tmpFile))
	os.Remove(tmpFile.Name())
}

func TestMd5Hex(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("/tmp", "godfs_test_1.tmp")
	ioutil.WriteFile(tmpFile.Name(), []byte("abc"), 600)
	println(Md5Hex(tmpFile))
	os.Remove(tmpFile.Name())
}

func TestCurrentDir(t *testing.T) {
	println(CurrentDir())
}

func TestNowDate(t *testing.T) {
	println(NowDateDir())
}
