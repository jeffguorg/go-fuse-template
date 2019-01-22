package internal

import (
	"context"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fuseutil"
)

type FileNode struct {
}

func (node *FileNode) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Mode = 0755
	attr.Size = 3
	return nil
}
func (f *FileNode) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	if !req.Flags.IsReadOnly() {
		return nil, fuse.Errno(syscall.EACCES)
	}
	resp.Flags |= fuse.OpenKeepCache
	return f, nil
}
func (f *FileNode) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {

	fuseutil.HandleRead(req, resp, []byte("123"))
	return nil
}
