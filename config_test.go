package godfs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	loadconf := LoadConfig("./godfs.yaml")
	bytes, _ := json.Marshal(loadconf)
	fmt.Println(string(bytes))
}

func TestGetConfigPath(t *testing.T) {
	_, e := GetConfigPath()
	fmt.Println(e)
}
