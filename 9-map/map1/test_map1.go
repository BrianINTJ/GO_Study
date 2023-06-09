package main

import (
	"fmt"
)

func main() {
	//声明myMap1是一种map类型 key是string， value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	//在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// ===> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "C"
	myMap2[2] = "go"
	myMap2[3] = "python"
	
	fmt.Println(myMap2)

	// ===> 第三种声明方式
	myMap3 := map[string]string{
		"one": "php",
		"two": "golang",
		"three": "python",
	}

	fmt.Println(myMap3)
}
