package main
import "fmt"
func main () {
var a,b,c float64
fmt.Print("Input 3 sides of a triangle ->")
fmt.Scan(&a,&b,&c)
if a*a+b*b>c*c {
	fmt.Println("The triangle is  obtuse")
}else{
	if a*a+b*b<c*c {
		fmt.Println("Yhe triagnle is acute")
	}else {
		fmt.Println("The triange is right")
}
}
}
