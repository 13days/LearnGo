package main

import "fmt"

func f1()  {
	// 使用make创建
	var city map[string]string
	city = make(map[string]string)

	city["中国"] = "北京"
	city["广东"] = "广州"
	city["河南"] = "郑州"

	mainCity, ok := city["中国"]
	if(ok){
		fmt.Printf("中国的首都:%s\n",mainCity)
	}


	// 删除
	delete(city, "中国")
	// 遍历
	for k,v := range city{
		fmt.Printf("{%s:%s}\n",k,v)
	}
}

func main() {
	f1()
}
