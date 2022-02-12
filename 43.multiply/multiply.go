package frytea_leetcode

import (
	// "fmt"
	// "reflect"
	// "strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2, tmp := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := len(n1) - 1; i >= 0; i-- {
		for j := len(n2) - 1; j >= 0; j-- {
			// 计算和
			tmp[i+j+1] += int((n1[i] - '0') * (n2[j] - '0'))
			// 进位
			if tmp[i+j+1] >= 10 {
				tmp[i+j] += tmp[i+j+1] / 10
				tmp[i+j+1] %= 10
				// fmt.Println(tmp)
			}
		}
	}

	// 去0
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}

	// fmt.Println(tmp)
	// 转 string
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = byte(tmp[i] + '0')
	}
	// fmt.Println(string(res))
	return string(res)

}
