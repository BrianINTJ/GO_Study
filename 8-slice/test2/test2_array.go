package main

import "fmt"

func printArray(myArray []int){
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	
	myArray[0] = 100
}

func main(){
	myArray1 := []int{1,2,3,4}	// 动态数组，切片 slice
	
	fmt.Printf("myArray type is %T\n", myArray1)
	printArray(myArray1);
	for _, value := range myArray1 {
		fmt.Println("value = ", value)
	}
}