package FindNativeMounts

import (
	"testing"
	"fmt"
	"sync"
)

func TestFindNativeMounts(t *testing.T) {
	wg := new(sync.WaitGroup)
	fileInfoChan := make(chan string, 10)
	wg.Add(1)
	go func() {
		if err := FindNativeMounts("/", fileInfoChan, wg); err != nil { // Can provide optional magic numbers
			panic(err)
		}
	}()
	for ; ; {
		info, ok := <- fileInfoChan
		if ok == true {
			fmt.Println(info, "is native")
		} else {
			fmt.Println("Finished")
			break
		}

	}
	wg.Wait()
}
