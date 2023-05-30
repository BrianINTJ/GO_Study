package main

import "fmt"

func printMap(cityMap map[string]string)  {
	//遍历
	//cityMap 是一个应用传递
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}
}

func ChangeValue(cityMap map[string]string)  {
	cityMap["England"] = "london"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	printMap(cityMap)
	
	//删除
	delete(cityMap, "China")
	
	//修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)
	
	fmt.Println("--------")
	
	//遍历
	printMap(cityMap)
	
}
