// DataType
package main

import (
	"fmt"
)

func main() {

	/*
	 * integer
	 * The default type of an integer is int.
	 * The rear number is the number of bytes of the corresponding type.
	 * The following types have their corresponding unsigned types. such as unit8...
	 */
	i := 0
	var i8 int8 = 1
	var i16 int16 = 2
	var i32 int32 = 3
	var i64 int64 = 4
	isum := int64(i8) + int64(i16) + int64(i32) + i64
	fmt.Printf("The type of i is %T\n", i)
	fmt.Printf("The type of i8 is %T\n", i8)
	fmt.Printf("The type of i16 is %T\n", i16)
	fmt.Printf("The type of i32 is %T\n", i32)
	fmt.Printf("The type of i64 is %T\n", i64)
	fmt.Printf("The type of isum is %T\n", isum)

	/*
	 * float
	 * The default type of an integer is float64.
	 * The rear number is the number of bytes of the corresponding type.
	 */
	f := 0.0
	var f32 float32 = 1.0
	var f64 float64 = 2.0
	fsum := float64(f32) + float64(f64)
	fmt.Printf("The type of f is %T\n", f)
	fmt.Printf("The type of f32 is %T\n", f32)
	fmt.Printf("The type of f64 is %T\n", f64)
	fmt.Printf("The type of fsum is %T\n", fsum)

}
