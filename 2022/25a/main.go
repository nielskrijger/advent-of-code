package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/25a/sample.txt")
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		total += convertSNAFUToDec(line)
	}
	fmt.Printf("Answer: %+v", convertDecToSNAFU(total))
}

var decode = map[uint8]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}

func convertSNAFUToDec(val string) int {
	r := 0
	for i := 0; i < len(val); i++ {
		r += decode[val[len(val)-i-1]] * int(math.Pow(5, float64(i)))
	}
	return r
}

var encode = map[int]string{0: "0", 1: "1", 2: "2", 3: "=", 4: "-"}

func convertDecToSNAFU(val int) string {
	var r string
	for val > 0 {
		r = encode[val%5] + r
		val = (val + 2) / 5
	}
	return r
}
