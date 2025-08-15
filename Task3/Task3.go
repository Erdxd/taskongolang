package Task3

import "fmt"

func Task3() {
	stats := []int{10, 5}
	add(stats, 2)
	fmt.Println(stats)

}
func add(stats []int, n int) {
	stats = append(stats, n)

}
