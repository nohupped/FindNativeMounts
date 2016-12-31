### FindNativeMounts

FindNativeMounts traverse the present working directory and finds if the directories
found are native to the filesystem or not, based on the filesystem magic number defined in constants.go taken
from /usr/include/linux/magic.h. This will not traverse recursively.
The function takes a directory path as a string, a string channel, a pointer to
sync.waitgroup and an optional Magic number constant(s) defined in constants.go.
If magic number is not specified, it assigns EXT4_SUPER_MAGIC by default that equals to 0xEF53.
If unsure of the magic number, use tune2fs to find it. This function don't handle
errors by itself, but returns it if occurs.

#### Install:

```
go get -u github.com/nohupped/FindNativeMounts
```

#### How to use:

```
package main
import (
	"fmt"
	"sync"
	"github.com/nohupped/FindNativeMounts"
)
func main() {
	wg := new(sync.WaitGroup)
	fileNamesChan := make(chan string, 10)
	wg.Add(1)
	go func() {
		if err := FindNativeMounts.Find("/", fileNamesChan, wg); err != nil { 
		// Can provide optional magic numbers, withoug which the function self-assigns ext4's magic number. 
		// Eg:
		//if err := FindNativeMounts.Find("/", fileNamesChan, wg, FindNativeMounts.NFS_SUPER_MAGIC, FindNativeMounts.PROC_SUPER_MAGIC); err != nil { 
		//This will show only the NFS mounts and the proc mounts.
			panic(err)
		}
	}()
	for ; ; {
		info, ok := <-fileNamesChan
		if ok == true {
			fmt.Println(info, "is native")
		} else {
			fmt.Println("Finished")
			break
		}

	}
	wg.Wait()
}
```

#### Output:

```
/lost+found is native
/bin is native
/initrd.img is native
/vmlinuz is native
/sbin is native
/etc is native
/lib64 is native
/lib is native
/home is native
/var is native
/tmp is native
/initrd.img.old is native
/nodejs is native
/Personal is native
/srv is native
/live-build is native
/usr is native
/root is native
/boot is native
/vmlinuz.old is native
/media is native
/opt is native
Finished
```

#### Original mounts:

```
devdesk# df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            3.4G     0  3.4G   0% /dev
tmpfs           695M  9.3M  686M   2% /run
/dev/sda7        51G  8.9G   39G  19% /
tmpfs           3.4G  390M  3.1G  12% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           3.4G     0  3.4G   0% /sys/fs/cgroup
/dev/sda3       492G  190G  277G  41% /Personal
/dev/sda1       4.5G  116M  4.1G   3% /boot
/dev/sda6        92G  5.3G   82G   7% /var
/dev/sda5       275G  182G   80G  70% /home
tmpfs           695M   28K  695M   1% /run/user/1000
localhost:/tmp   51G  8.9G   39G  19% /mnt
```
