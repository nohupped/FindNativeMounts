package FindNativeMounts

import (
	"os"
	"syscall"
	"sync"
	"bytes"
)



// FindNativeMounts traverse the present working directory and finds if the directories
// found are native to the filesystem or not, based on the filesystem magic number defined in constants.go taken
// from /usr/include/linux/magic.h. This will not traverse recursively.
// The function takes a directory path as a string, a string channel, a pointer to
// sync.waitgroup and an optional Magic number constant(s) defined in constants.go.
// If magic number is not specified, it assigns EXT4_SUPER_MAGIC by default that equals to 0xEF53.
// If unsure of the magic number, use tune2fs to find it. This function don't handle
// errors by itself, but returns it if occurs.
func FindNativeMounts(dir string, fileInfoChan chan string, wg *sync.WaitGroup, nativePartitions ...int64) error{
	if len(nativePartitions) <= 0 {
		nativePartitions = []int64{EXT4_SUPER_MAGIC}
	}
	x, err := os.Open(dir)
	if err != nil {
		return err
	}
	y, err := x.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, i := range y {
		var buff bytes.Buffer
		buff.WriteString(dir)
		switch dir {
		case "/":
		default:
			buff.WriteString("/")
		}

		buff.WriteString(i)
		/*err := os.Chdir(dir)
		if err != nil {
			return err
		}*/

		t := new(syscall.Statfs_t)
		err = syscall.Statfs(buff.String(), t)
		if err != nil {
			return err
		}
		if checkDirIsNative(t.Type, nativePartitions) {
			//fmt.Println(i.Name(), "is native")
			fileInfoChan <- buff.String()
		} else {
			//fmt.Println(i.Name(), "is not native")
		}
	}
	close(fileInfoChan)
	wg.Done()
	return nil
}

func checkDirIsNative(dirtype int64, nativetypes []int64) bool{
	for _, i := range nativetypes {
		if dirtype == i {
			return true
		}
	}
	return false
}
