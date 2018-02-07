// Levenshtein Distance in Golang
package main

import "fmt"

// 在信息论和计算机科学中，莱文斯坦距离是一种两个字符串序列的距离度量。形式化地说，两个单词的莱文斯坦距离是一个单词变成另一个单词要求的最少单个字符编辑数量（如：删除、插入和替换）。莱文斯坦距离也被称做编辑距离，尽管它只是编辑距离的一种，与成对字符串比对紧密相关。
func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func main() {
	var str1 = []rune("Asheville")
	var str2 = []rune("Arizona")
	fmt.Println("Distance between Asheville and Arizona:", levenshtein(str1, str2))

	str1 = []rune("Python")
	str2 = []rune("Peithen")
	fmt.Println("Distance between Python and Peithen:", levenshtein(str1, str2))

	str1 = []rune("cui")
	str2 = []rune("hao")
	fmt.Println("Distance between cui and hao:", levenshtein(str1, str2))
}
