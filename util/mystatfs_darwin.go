package util

//#include <stdlib.h>
//#include <sys/param.h>
//#include <sys/mount.h>
//typedef struct statfs statfs_t;
import "C"
import (
	"fmt"
	"unsafe"
)

func GetStatus(myStatFs *MyStatFs, path string) error {

	var darwinStatFs C.statfs_t

	CPath := C.CString(path)
	defer C.free(unsafe.Pointer(CPath))
	err2 := C.statfs(CPath, &darwinStatFs)
	if err2 != 0 {
		return fmt.Errorf("statfs error %d", err2)
	}

	// Available blocks * size per block = available space in bytes
	myStatFs.Avail = uint64(darwinStatFs.f_bavail) * uint64(darwinStatFs.f_bsize)
	myStatFs.Total = uint64(darwinStatFs.f_blocks) * uint64(darwinStatFs.f_bsize)

	return nil
}
