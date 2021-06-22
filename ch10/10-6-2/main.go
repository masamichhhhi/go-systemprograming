package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/billziss-gh/cgofuse/fuse"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

type CloudFileSystem struct {
	fuse.FileSystemBase
	bucket *blob.Bucket
}

func (cf *CloudFileSystem) Getattr(path string, stat *fuse.Stat_t, fh uint64) (errc int) {
	if path == "/" {
		stat.Mode = fuse.S_IFDIR | 0555
		return 0
	}
	ctx := context.Background()
	name := strings.TrimLeft(path, "/")
	a, err := cf.bucket.Attributes(ctx, name)
	if err != nil {
		_, err := cf.bucket.Attributes(ctx, name+"/")
		if err != nil {
			return fuse.ENOENT
		}
		stat.Mode = fuse.S_IFDIR | 0555
	} else {
		stat.Mode = fuse.S_IFDIR | 0444
		stat.Size = a.Size
		stat.Mtim = fuse.NewTimespec(a.ModTime)
	}
	stat.Nlink = 1
	return 0
}

func main() {
	ctx := context.Background()
	if len(os.Args) < 3 {
		fmt.Printf("%s [bucket-path] [mount-point] etc...", os.Args[0])
		os.Exit(1)
	}
	b, err := blob.OpenBucket(ctx, os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b.Close()
	cf := &CloudFileSystem{bucket: b}
	host := fuse.NewFileSystemHost(cf)
	host.Mount(os.Args[2], os.Args[3:])
}
