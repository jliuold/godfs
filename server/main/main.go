package main

import "github.com/godfs-io/godfs/server"

func main() {
	server.StartHttp2()
	select {}
}
