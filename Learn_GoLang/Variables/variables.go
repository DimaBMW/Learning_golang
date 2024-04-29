package main
// here I will talk about variables in golang
func main(){
	//to create a variable you need to use the var keyword
	//for example
	var Name string // string this is type variable

	// types variable
	var a int // type integer default your system encoding
	var b int8 // type integer value 8 bits => range of values [-128 : 127]
	var c int16 // type integer value 16 bits => range of values  [-32768 : 32767]
	var d int32 // type integer value 32 bits => range of values [-2147483648 : 2147483647]
	var e int64 // type integer value 64 bits => range of values  [-9223372036854775808 : 9223372036854775807]

	var f uint // type integer non-negative default your system encoding
	var g uint8 // type integer non-negative value 8 bits => range of values [0 : 127]
	var h uint16 // type integer non-negative value 16 bits => range of values [0 : 32767]
	var j uint32 // type integer non-negative value 32 bits => range of values [0 : 2147483647]
	var k uint64 // type integer non-negative value 64 bits => range of values [0 : 9223372036854775807]

	var str string // string data type ('')
	
	var true bool // boolean data type (true or false / 0 or 1)

	var ptr uintptr // Go's uintptr type is an unsigned integer type large enough to store the uninterpreted bits of a pointer value.

	var num float32 // float32 is a single precision floating point number that can represent values ​​up to 7 decimal digits.
	var num2 float64 // float64 is a single precision floating point number that can represent values ​​up to 15 decimal digits.

	var compl complex64 //complex64 is a complex number with a real and an imaginary part, each represented as a single-precision floating point number. (1.23+4.56i)
	var compl2 complex128 //complex128 is a complex number with a real and an imaginary part, each represented as a single-precision floating point number.(1.23+4.56i)

	var byte byte // In Go, a bytetype is an alias for the uint8 type, which is an unsigned 8-bit integer. bs := []byte("◺")

	var rune rune //rune in Go is a 32-bit value that represents a Unicode code point. This is an alias to int32 which is used to handle individual characters in a string

	//short variable declaration
	name_2 := "Dima" // type of int
	symbol := 'a' // type of rune
	flag := false // type of bool

	//assigning other values ​​to variables
	var doc int8 = 2
	doc = 32
	println(doc) // 32

	//actions with variables
	/**
		+ // addition
		- // subtraction
		* // multiplication
		% // remainder of the division

		// bitwise binary arithmetic operators

		& // logical and
		| // Boolean or
		^ // Boolean works like this: it returns true if exactly one of its operands is true, and false otherwise
		&^ // bit clearing
		! // logical not
		<< // shift left by bit
		>> // shift right by bit
		++ // increment +1
		-- // decrement -1
		>= // greater than or equal to
		<= // less than or equal to
		== // equals
	*/ 
}