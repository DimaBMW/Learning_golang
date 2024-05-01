package main

func main(){
	//constants constant variable - value cannot be changed
	//constant declaration
	const π = 3.1416
	// or
	const (
		A, B int64   = -3, 5
		Y    float32 = 2.718
 	)

	//iota in golang
	// ioat automatically assigns values ​​to constants in order
	const (
		a = iota
		b = iota
		c = iota
	)
	println(a,b,c) // 0,1,2
}