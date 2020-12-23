package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64 (number of decimal places)
	// uint uint8 uint16 uint32 uint64 uintptr (unsigned integer - 0+ or => 0)
	// byte - alias for int8
	// rune - alias for int32
	// float32 float64
	// complex32 complex128 (really large numbers)

	// Using var and const
	// var name string = "Gran" // do not need to put string
	var ageInt32 int32 = 14
	var ageInt int = 14
	var isCool = true // if changed from 'var' to 'const'
	isCool = false    // this line will not work, and an error "./main.go:22:9: cannot assign to isCool"
	var size float32 = 2.3

	// Shorthand
	// name := "Granwyn"
	// size := 1.3
	// email := "email@domainname.domainsuffix"

	name, email := "Granwyn", "email@domainname.domainsuffix"

	fmt.Println(name, ageInt, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", ageInt32)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
