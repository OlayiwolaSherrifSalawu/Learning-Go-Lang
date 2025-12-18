package layi

import "fmt"

func Union(s, s2 string) {
	result := ""
	newStr := s + s2
	if len(s) == 0 {
		fmt.Println(s)
	}
	for i := 0; i < len(newStr); i++ {
		isSame := true
		for j := 0; j < len(result); j++ {
			if newStr[i] == result[j] {
				isSame = false
			}
		}
		if isSame {
			result += string(newStr[i])
		}
	}
	fmt.Println(result)
}
