package internal

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type DirNode struct {
	file string
}

var dirDirs = []fuse.Dirent{
	{Inode: 2, Name: "clock", Type: fuse.DT_File},
}

func (node *DirNode) Attr(ctx context.Context, attr *fuse.Attr) error {
	logrus.Infoln("dir attr")
	attr.Mode = os.ModeDir | 0755
	return nil
}

func (d *DirNode) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == "clock" {
		return &FileNode{}, nil
	}
	return nil, fuse.ENOENT
}
func (d *DirNode) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return dirDirs, nil
}
