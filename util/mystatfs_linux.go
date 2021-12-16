package util

import (
	"golang.org/x/sys/unix"
)

func GetStatus(stat *MyStatFs, path string) error {

	var unixStatFs unix.Statfs_t

	err2 := unix.Statfs(path, &unixStatFs)
	if err2 != nil {
		panic(err2)
	}

	// Available blocks * size per block = available space in bytes
	stat.Avail = unixStatFs.Bavail * uint64(unixStatFs.Bsize)
	stat.Total = unixStatFs.Blocks * uint64(unixStatFs.Bsize)

	return nil
}
