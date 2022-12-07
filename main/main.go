package main

import (
	"bufio"
	"os"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	m := map[int]int{}

	for _, elem := range str {
		if value, inMap := m[int(elem)]; inMap {
			continue
		} else {
			m[elem] = work(int(elem))
		}
	}

}
