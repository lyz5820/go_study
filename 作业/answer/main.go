package main

import (
	"fmt"
)

func myBase64Encode(input string) string {
	const table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	data := []byte(input)
	var result string
	n := len(data)
	for i := 0; i < n; i += 3 {
		var b1, b2, b3 int
		b1 = int(data[i])
		val := b1 << 16
		if i+1 < n {
			b2 = int(data[i+1])
			val |= b2 << 8
		}
		if i+2 < n {
			b3 = int(data[i+2])
			val |= b3
		}
		result += string(table[(val>>18)&0x3F])
		result += string(table[(val>>12)&0x3F])
		if i+1 < n {
			result += string(table[(val>>6)&0x3F])
		} else {
			result += "="
		}
		if i+2 < n {
			result += string(table[val&0x3F])
		} else {
			result += "="
		}
	}
	return result
}

func myBase64Decode(input string) string {
	const table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// 建立反向索引表：字符 -> 索引
	revTable := make([]int, 256)
	for i := 0; i < len(table); i++ {
		revTable[table[i]] = i
	}

	var result []byte
	// 每4个字符为一组进行处理
	for i := 0; i < len(input); i += 4 {
		if input[i] == '=' {
			break
		}

		// 取出4个6位索引
		v1 := revTable[input[i]]
		v2 := revTable[input[i+1]]
		v3, v4 := 0, 0

		hasV3, hasV4 := false, false
		if input[i+2] != '=' {
			v3 = revTable[input[i+2]]
			hasV3 = true
		}
		if input[i+3] != '=' {
			v4 = revTable[input[i+3]]
			hasV4 = true
		}

		// 拼接成24位整数
		val := (v1 << 18) | (v2 << 12) | (v3 << 6) | v4

		// 还原成字节
		result = append(result, byte((val>>16)&0xFF))
		if hasV3 {
			result = append(result, byte((val>>8)&0xFF))
		}
		if hasV4 {
			result = append(result, byte(val&0xFF))
		}
	}
	return string(result)
}

func myHexEncode(input string) string {
	const hexChars = "0123456789abcdef"
	var result string
	for i := 0; i < len(input); i++ {
		result += string(hexChars[input[i]>>4])
		result += string(hexChars[input[i]&0x0F])
	}
	return result
}

func myHexDecode(input string) string {
	const hexChars = "0123456789abcdef"
	var result []byte
	for i := 0; i < len(input); i += 2 {
		var high, low int
		for j := 0; j < 16; j++ {
			if input[i] == hexChars[j] {
				high = j
			}
			if input[i+1] == hexChars[j] {
				low = j
			}
		}
		result = append(result, byte(high<<4|low))
	}
	return string(result)
}

func myUrlEncode(input string) string {
	const hexChars = "0123456789ABCDEF"
	var result string
	for i := 0; i < len(input); i++ {
		c := input[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			result += string(c)
		} else {
			result += "%" + string(hexChars[c>>4]) + string(hexChars[c&0x0F])
		}
	}
	return result
}

func myUrlDecode(input string) string {
	const hexChars = "0123456789ABCDEF"
	const hexCharsLower = "0123456789abcdef"
	var result []byte
	for i := 0; i < len(input); i++ {
		if input[i] == '%' && i+2 < len(input) {
			var high, low int
			for j := 0; j < 16; j++ {
				if input[i+1] == hexChars[j] || input[i+1] == hexCharsLower[j] {
					high = j
				}
				if input[i+2] == hexChars[j] || input[i+2] == hexCharsLower[j] {
					low = j
				}
			}
			result = append(result, byte(high<<4|low))
			i += 2
		} else {
			result = append(result, input[i])
		}
	}
	return string(result)
}

func main() {
	for {
		fmt.Println("编码输入1，解码输入2，退出输入exit")
		var choice string
		fmt.Scanln(&choice)
		if choice == "exit" {
			break
		}

		for {
			mode := "编码"
			if choice == "2" {
				mode = "解码"
			}
			fmt.Printf("base64%s输入1，url%s输入2，hex%s输入3，exit返回主页面\n", mode, mode, mode)
			var sub string
			fmt.Scanln(&sub)
			if sub == "exit" {
				break
			}

			fmt.Printf("输入你要%s的字符串：\n", mode)
			var txt string
			fmt.Scanln(&txt)

			if choice == "1" {
				switch sub {
				case "1":
					fmt.Printf("编码后的内容为： %s\n", myBase64Encode(txt))
				case "2":
					fmt.Printf("编码后的内容为： %s\n", myUrlEncode(txt))
				case "3":
					fmt.Printf("编码后的内容为： %s\n", myHexEncode(txt))
				}
			} else {
				switch sub {
				case "1":
					fmt.Printf("解码后的内容为： %s\n", myBase64Decode(txt))
				case "2":
					fmt.Printf("解码后的内容为： %s\n", myUrlDecode(txt))
				case "3":
					fmt.Printf("解码后的内容为： %s\n", myHexDecode(txt))
				}
			}
		}
	}
}
