package main

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

func main() {
	cfs := pathfs.NewPathNodeFs(&ChatFS{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot("/tmp/chatfs", cfs.Root(), nil)
	if err != nil {
		panic(err)
	}
	server.Serve()
}
