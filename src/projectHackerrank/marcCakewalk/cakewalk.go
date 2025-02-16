// Package cakewalk : https://www.hackerrank.com/challenges/marcs-cakewalk
package cakewalk

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Cakewalk soln
func Cakewalk() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	list := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Fscan(io, &list[j])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	miles := 0
	for i := 0; i < n; i++ {
		miles += int(math.Pow(2, float64(i))) * list[i]
	}
	fmt.Println(miles)

}
