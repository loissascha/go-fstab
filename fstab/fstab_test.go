package fstab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringParsing(t *testing.T) {
	entries, err := parseStr(`
# Static information about the filesystems.
# See fstab(5) for details.

# <file system> <dir> <type> <options> <dump> <pass>
# /dev/nvme2n1p4
UUID=7a235ef76a	/         	ext4      	rw,relatime,stripe=4	0 1

# /dev/nvme2n1p2
UUID=9e6757da	/boot     	ext4      	rw,relatime,stripe=32	0 2

# /dev/nvme2n1p1
UUID=E90A-1DBE      	/boot/efi 	vfat      	rw,relatime,fmask=0022,dmask=0022,codepage=437,iocharset=ascii,shortname=mixed,utf8,errors=remount-ro	0 2

UUID=36614       /home   ext4    defaults,noatime       0       2

  UUID=dbe746 /mnt/disk2 ext4 defaults,noatime 0 2  

//192.168.0.100/Media /mnt/media cifs credentials=/etc/samba/creds/my_server_creds,uid=1000,gid=1000,file_mode=0664,dir_mode=0775,iocharset=utf8,vers=3.0,nofail,soft,noserverino,x-systemd.automount,_netdev	0	0
	`)

	assert.Nil(t, err)
	assert.Equal(t, 6, len(entries))

	rootFs := entries[0]
	assert.Equal(t, "UUID=7a235ef76a", rootFs.Device)
	assert.Equal(t, "/", rootFs.Mountpoint)
	assert.Equal(t, "ext4", rootFs.FsType)
	assert.Equal(t, 3, len(rootFs.Options))
	assert.Equal(t, "rw", rootFs.Options[0])
	assert.Equal(t, "relatime", rootFs.Options[1])
	assert.Equal(t, "stripe=4", rootFs.Options[2])
	assert.Equal(t, "0", rootFs.Dump)
	assert.Equal(t, "1", rootFs.Fsck)

	assert.Equal(t, "UUID=dbe746", entries[4].Device)
	assert.Equal(t, "/mnt/disk2", entries[4].Mountpoint)
}
