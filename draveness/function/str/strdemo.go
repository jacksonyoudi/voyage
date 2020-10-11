package main

import "fmt"

/**
1. 长度 len

golang的编码统一是utf8， ascII占一个字节，汉字占用3个字节

遍历字符串
2.  []rune(str)

字符串转数字
3. strconv.Atoi(str)

数字转字符串
4. strconv.Itoa()


5. 字符串转[]byte
[]byte(str)

6. []byte 转成字符串
str = string([]byte{97,98,99})

7. 10进制转 2，8 16 进制
strconv.FormatInt(int64(i), 16)

8. strings.Contains(str, s)

9. strings.Count(str, s)

10. 比较不区分大小写
strings.EqualFold("abc", "ABC")

11. strings.Index(s, str)

12. strings.LastIndex(str, s)


13. strings.Replace(str, s)

14.

*/
func main() {
	i := len("a")
	fmt.Println(i)

}
