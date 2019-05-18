package main

import ("fmt" 
		"math"
		"math/rand") 

func demo(){
	// var num1 float64 = 4.4
	// var num2 float64 =9.6
	// var num3,num4 float64 =2.3 ,4.5
	// num5,num6 :=3.3,4.4
	// fmt.Println(TypesDemo(num1,num2))
	// fmt.Println(TypesDemo(5.5,4.5))
	// fmt.Println(num3+num4)
	// fmt.Println(num5+num6)

	w1,w2 :="hi "," threre"
	fmt.Println(multiple(w1,w2))
}
func multiple(a,b string)(string,string){
	return a,b
}
func foo(){

	fmt.Println("Square root of 4 is",math.Sqrt(4))
}
func randomnumber(){

	fmt.Println("Random number between 1-100",rand.Intn(100))
	fmt.Println("Random number between 1-100",rand.Intn(100))
}
func TypesDemo(x float64,y float64) float64{
	return x+y
}
func Parsing(){
	// var a int =62
	// var b float64=float64(a)

	// x:=a
}