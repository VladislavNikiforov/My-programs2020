package main
import "fmt"
import "math"
func main () {
var a,b,c,D,x1,x2 float64
fmt.Println("Input the numbers to quadratic equal")
fmt.Scan(&a,&b,&c)
D=b*b-4*a*c
if D>0 {
	x1=(-b-math.Sqrt(D))/(2*a)
	x2=(-b+math.Sqrt(D))/(2*a)
	fmt.Printf("The roots are %f and %f",x1,x2)
} else {
	if D==0 {
	x1=-b/2*a
	fmt.Printf("the root is equal to %f",x1)
	} else {
		fmt.Println("There are not any roots for this equal")

}
}
}
                                     
