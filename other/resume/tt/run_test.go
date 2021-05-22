package tt

import (
	"fmt"
	"testing"
)

//func TestRun(t *testing.T) {
//	run()
//}

func TestRunF(t *testing.T) {
	//funcf()
	var flist []func()
	for i := 0; i < 3; i++ {
		i := i
		flist = append(flist, func() {
			fmt.Printf("%p", &i)
			// 3 3 3
			println(i)
		})
		for j := 0; j < 3; j++ {
			fmt.Printf("%p", flist[j])
		}
	}
}
