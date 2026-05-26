package fstab

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.NoError(t, err)
	assert.Equal(t, 6, len(entries))

	rootFs := entries[0]
	assert.Equal(t, "UUID=7a235ef76a", rootFs.Device)
	assert.Equal(t, "/", rootFs.Mountpoint)
	assert.Equal(t, "ext4", rootFs.FsType)
	assert.Equal(t, []string{"rw", "relatime", "stripe=4"}, rootFs.Options)
	assert.Equal(t, "0", rootFs.Dump)
	assert.Equal(t, "1", rootFs.Fsck)

	bootFs := entries[1]
	assert.Equal(t, "UUID=9e6757da", bootFs.Device)
	assert.Equal(t, "/boot", bootFs.Mountpoint)
	assert.Equal(t, "ext4", bootFs.FsType)
	assert.Equal(t, []string{"rw", "relatime", "stripe=32"}, bootFs.Options)
	assert.Equal(t, "0", bootFs.Dump)
	assert.Equal(t, "2", bootFs.Fsck)

	efiFs := entries[2]
	assert.Equal(t, "UUID=E90A-1DBE", efiFs.Device)
	assert.Equal(t, "/boot/efi", efiFs.Mountpoint)
	assert.Equal(t, "vfat", efiFs.FsType)
	assert.Equal(t, []string{"rw", "relatime", "fmask=0022", "dmask=0022", "codepage=437", "iocharset=ascii", "shortname=mixed", "utf8", "errors=remount-ro"}, efiFs.Options)
	assert.Equal(t, "0", efiFs.Dump)
	assert.Equal(t, "2", efiFs.Fsck)

	homeFs := entries[3]
	assert.Equal(t, "UUID=36614", homeFs.Device)
	assert.Equal(t, "/home", homeFs.Mountpoint)
	assert.Equal(t, "ext4", homeFs.FsType)
	assert.Equal(t, []string{"defaults", "noatime"}, homeFs.Options)
	assert.Equal(t, "0", homeFs.Dump)
	assert.Equal(t, "2", homeFs.Fsck)

	disk2Fs := entries[4]
	assert.Equal(t, "UUID=dbe746", disk2Fs.Device)
	assert.Equal(t, "/mnt/disk2", disk2Fs.Mountpoint)
	assert.Equal(t, "ext4", disk2Fs.FsType)
	assert.Equal(t, []string{"defaults", "noatime"}, disk2Fs.Options)
	assert.Equal(t, "0", disk2Fs.Dump)
	assert.Equal(t, "2", disk2Fs.Fsck)

	mediaFs := entries[5]
	assert.Equal(t, "//192.168.0.100/Media", mediaFs.Device)
	assert.Equal(t, "/mnt/media", mediaFs.Mountpoint)
	assert.Equal(t, "cifs", mediaFs.FsType)
	assert.Equal(t, []string{"credentials=/etc/samba/creds/my_server_creds", "uid=1000", "gid=1000", "file_mode=0664", "dir_mode=0775", "iocharset=utf8", "vers=3.0", "nofail", "soft", "noserverino", "x-systemd.automount", "_netdev"}, mediaFs.Options)
	assert.Equal(t, "0", mediaFs.Dump)
	assert.Equal(t, "0", mediaFs.Fsck)
}
