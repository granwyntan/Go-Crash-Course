package main

import (
	"fmt"
	"math"

	"github.com/granwyntan/go_crash_course/03_Packages/strutil"
	//"strutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
	fmt.Println(strutil.Reverse("hello"))
}
