package main

import (
	"fmt"
	"strconv"
)

func binaryToDecimal(binary string) int {
	dec, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func decimalToBinary(dec int) string {
	return fmt.Sprintf("%b", dec)
}

func main() {
	fmt.Println(binaryToDecimal("1001"))
	fmt.Println(decimalToBinary(19))
}
