package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 基数排序的时间复杂度是 O(k*n)，其中  n是排序元素个数，  k是数字位数。注意这不是说这个时间复杂度一定优于 O(n*log(n)), k的大小取决于数字位的选择（比如比特位数），和待排序数据所属数据类型的全集的大小； k决定了进行多少轮处理，而 n是每轮处理的操作数目。
func main() {
	us := []int32{6, 1, 3, 9, 2, 7, 4, 8, 5}
	radixsort(us)
	fmt.Println(us)
}

const digit = 4
const maxbit = -1 << 31

func radixsort(s []int32) {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(s))
	for i, e := range s {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		s[i] = w ^ maxbit
	}
}
