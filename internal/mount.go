package internal

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// Mount a path to
func Mount(path, mountpoint string) error {

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		return err
	}
	defer c.Close()

	filesys := &FS{}
	if err := fs.Serve(c, filesys); err != nil {
		return err
	}

	// check if the mount process has an error to report
	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}

	return nil
}
