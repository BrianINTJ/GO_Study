package main

// import _ "fmt"	给fmt包起一个别名，匿名，无法使用当前包的方法，但是会执行当前的包的内部的init()方法
// import aa "fmt"	给包起一个别名，aa，aa.Println()来直接调用
// import . "fmt"	将当前fmt包中的全部方法，导入到当前本包的作用中，fmt包中的全部的方法可以直接使用API来调用，不需要fmt.API来调用 （不建议使用）
import (
	_ "GoStudy/5-init/lib1"
	mylib2 "GoStudy/5-init/lib2"
	// . "GoStudy/5-init/lib2"
)

func main() {
	// lib1.Lib1Test()
	mylib2.Lib2Test()
	// Lib2Test()
}
