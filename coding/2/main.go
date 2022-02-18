package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(kalimatPalindromeTerpanjang("aku suka makan nasi"))
	fmt.Println(kalimatPalindromeTerpanjang("di rumah saya ada kasur rusak"))
	fmt.Println(kalimatPalindromeTerpanjang("abcde edcbza"))
	// a := "abc"
	// fmt.Println(string(a[0]))
}

func kalimatPalindromeTerpanjang(k string) (output string) {
	var strings1 = make([]string, 0)
	// var stringsRev1 = make([]string, len(k))

	for _, v := range k {
		strings1 = append(strings1, string(v))
	}
	// for i := len(k) - 1; i >= 0; i-- {
	// 	stringsRev1[i] = string(k[i])
	// }

	var outputSlc = make([]string, 0)
	

	// var temp1 = strings.Join(strings1, "")
	// var temp2 = strings.Join(stringsRev1, "")
	// if strings.Contains(temp1, temp2) {
		// temp3 := []string{}
		// for i := len(temp2) - 1; i >= 0; i-- {
		// 	temp3 = append(temp3, string(temp2[i]))
		// }
		// temp4 := []string{}
		// loop:
		// 	for _, v := range temp1 {
		// 		val1 := string(v)
		// 		for _, v := range temp3 {
		// 			if v != val1 {
		// 				continue loop
		// 			}
		// 			outputSlc = append(outputSlc, string(v))
		// 			break
		// 		}
		// 	}
		// for i := 0; i < len(temp1); i++ {
		// 	val := string(temp1[i])
		// 	for i := len(temp2) - 1; i >= 0; i-- {
		// 		if string(temp2[i]) == val {
		// 			outputSlc = append(outputSlc, val)
		// 		}
		// 	}
		// }
		// return strings.Join(outputSlc, "")
	}
	return ""
}
