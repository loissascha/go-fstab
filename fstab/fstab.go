package fstab

type FsEntry struct {
	Device     string
	Mountpoint string
	FsType     string
	Options    string
	Dump       string
	Fsck       string
}

func ReadFile(path string) ([]FsEntry, error) {
	return nil, nil
}
