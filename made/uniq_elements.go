package main

import (
	"fmt"
)

func main() {
	cnt := make(map[int]int)
	var n, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		_, exists := cnt[x]
		if exists {
			cnt[x] += 1
		} else {
			cnt[x] = 1
		}
	}
	ans := 0
	for _, value := range cnt {
		if value == 1 {
			ans += 1
		}
	}
	fmt.Println(ans)
}
