package main 
import "fmt"
import "math"
func main() {
var h,min,Ah,Amin,Angle,x float64
fmt.Print("Hours -> ")
fmt.Scan(&h)
if (h>=12)  {
	h=h-12
}
fmt.Print("Minutes -> ")
fmt.Scan(&min)

Amin=6*min
Ah=30*h+0.5*min
Angle=Ah-Amin
if Angle>360 { 
	fmt.Println("Time was incorrectly inputed")
} else {
x=math.Abs(Angle)
}
fmt.Print("The angle is ", x )
}

