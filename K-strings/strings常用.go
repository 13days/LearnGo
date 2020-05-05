package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "1234412353dsfds"
	// 是否包含整串
	fmt.Println(strings.Contains(a, "1232"))
	// 是否包含任意一个
	fmt.Println(strings.ContainsAny(a,"1232"))
	// 字符串统计
	fmt.Println(strings.Count(a, "123"))
	// 前后缀
	fmt.Println(strings.HasPrefix(a,"123"))
	fmt.Println(strings.HasPrefix(a,"sdfs"))
	fmt.Println(strings.HasSuffix(a,"ds"))
	// 首次下标出现位置,不存在返回-1
	fmt.Println(strings.Index(a,"23"))
	// 任意字符第一次出现位置,
	fmt.Println(strings.IndexAny(a, "ls3"))
	// 查找最后一次出现的
	fmt.Println(strings.LastIndex(a, "2353dsfds"))
	// 拼接
	ss := []string{"SE", "171549333","吴楚鹏"}
	s := strings.Join(ss,"-")
	fmt.Println(s)
	// 切割
	t := strings.Split(s, "-")
	fmt.Println(t)
	// 重复拼接
	s = strings.Repeat("A",9)
	fmt.Println(s)
	// 替换,-1表示所有
	s = strings.Replace(s, "AA", "BB", 3)
	fmt.Println(s)
	s = strings.Replace(s, "BBB", "XXX", -1)
	fmt.Println(s)
	// 大小写
	s = "hello,world"
	s = strings.ToUpper(s)
	fmt.Println(s)
	s = strings.ToLower(s)
	fmt.Println(s)
	// 截取
	s = s[4:9]
	fmt.Println(s)
}
