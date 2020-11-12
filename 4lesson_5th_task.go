package main
import "fmt"
import "math"
//не работает, так как в итоге остается одно число и 3 нуля
func main () {
	var a1,a2,a3,a4 float64
	fmt.Println("Input 4 numbers")
	fmt.Scan(&a1,&a2,&a3,&a4)
		for true {
			a1=math.Abs(a2-a1)
			a2=math.Abs(a3-a2)
			a3=math.Abs(a4-a3)
			a4=math.Abs(a1-a4)
			fmt.Println(a1)
			fmt.Println(a2)
			fmt.Println(a3)
			fmt.Println(a4)
			if a1==0 && a2==0 && a3==0 && a4==0 {
				break 
			}
		}
		fmt.Println()
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)		

}
