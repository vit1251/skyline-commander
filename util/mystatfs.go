package util

type MyStatFs struct {
	Type       int
	TypeName   string
	MountPoint string
	MountRoot  string
	Device     string
	Avail      uint64 /* in kB */
	Total      uint64 /* in kB */
	//	nfree uint64
	//	nodes uint64
}
