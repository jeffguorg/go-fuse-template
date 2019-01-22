package internal

import (
	"bazil.org/fuse/fs"
	"github.com/sirupsen/logrus"
)

type FS struct{}

func (fs *FS) Root() (fs.Node, error) {
	logrus.Infoln("fs root")
	return &DirNode{}, nil
}
