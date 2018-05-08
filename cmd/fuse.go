// Fuse structure and implementation
package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/Wheeeel/chatfs/model"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type ChatFS struct {
	pathfs.FileSystem
}

const (
	LEVEL_ROOT   = 1
	LEVEL_SERVER = 2
	LEVEL_CHAT   = 3
)

func (fs *ChatFS) SetDebug(debug bool) {}

// ChatFS
func (fs *ChatFS) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	log.Info("CALL GetAttr")
	println("GetAttr:", name)
	attr := new(fuse.Attr)
	level := GetDirLevel(name)
	if name == "" {
		attr.Mode = fuse.S_IFDIR | 0755
	}
	if level == LEVEL_SERVER {
		attr.Mode = fuse.S_IFDIR | 0755
	}
	return attr, fuse.OK
}

func (fs *ChatFS) Readlink(name string, context *fuse.Context) (string, fuse.Status) {
	log.Info("CALL Readlink")
	return "", fuse.ENOSYS
}

func (fs *ChatFS) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	// TODO: get the level of the dir
	log.Info("CALL Mkdir")
	log.Info("Enter mkdir")
	level := GetDirLevel(name)
	if level == LEVEL_ROOT {
		// add the server
		model.AddServer(model.Server{
			Name:     name,
			Type:     "",
			User:     "",
			Password: "",
		})
	}
	if level == LEVEL_SERVER {

	}
	if level == LEVEL_CHAT {

	}
	return fuse.OK
}

func (fs *ChatFS) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Unlink")
	return fuse.ENOSYS
}

func (fs *ChatFS) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Rmdir")
	return fuse.ENOSYS
}

func (fs *ChatFS) Symlink(value string, linkName string, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Symlink")
	return fuse.ENOSYS
}

func (fs *ChatFS) Rename(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Rename")
	return fuse.ENOSYS
}

func (fs *ChatFS) Link(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Link")
	return fuse.ENOSYS
}

func (fs *ChatFS) Chmod(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Chmod")
	return fuse.ENOSYS
}

func (fs *ChatFS) Chown(name string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Chown")
	return fuse.ENOSYS
}

func (fs *ChatFS) Truncate(name string, offset uint64, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Truncate")
	return fuse.ENOSYS
}

func (fs *ChatFS) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	log.Info("CALL Open")
	return nil, fuse.ENOSYS
}

func (fs *ChatFS) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
	log.Info("CALL OpenDir", name)
	// 1. get a list of server opened
	if name == "" {
		slist, err := model.GetServers()
		if err != nil {
			log.Error(err)
		}

		for _, s := range slist {
			ent := fuse.DirEntry{}
			ent.Name = s.Name
			ent.Mode = fuse.S_IFREG

			stream = append(stream, ent)
		}
	}
	return stream, fuse.OK
}

func (fs *ChatFS) OnMount(nodeFs *pathfs.PathNodeFs) {
	// TODO: Initialize the hierarchy previously created

	// TODO: Perform the login process (If login failed, mount denied)

	// 1. select * from servers

	// 2. login server

	// 3. select * from channels

	// 4. init channel dir

	// 5. done
}

func (fs *ChatFS) OnUnmount() {

}

func (fs *ChatFS) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	log.Info("CALL Access")
	return fuse.ENOSYS
}

func (fs *ChatFS) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	log.Info("CALL Create")
	// switch
	// case server
	// -> how to login?
	// <- create a file named nickname, store your nickname
	// -> -> nickname is taken, is registerd [IGNORE]

	// ->

	// case channel
	return nil, fuse.ENOSYS
}

func (fs *ChatFS) Utimens(name string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *ChatFS) String() string {
	return "ChatFS"
}

func (fs *ChatFS) StatFs(name string) *fuse.StatfsOut {
	return nil
}
