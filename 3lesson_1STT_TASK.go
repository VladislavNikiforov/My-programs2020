package main 
import "fmt"
import "math"

func main () {
	var x1,x2,y1,y2,r1,r2 float64
	fmt.Println("Input coordinates of two centres (x1,y1; x2,y2) ->")
	fmt.Scan(&x1,&y1,&x2,&y2)
	fmt.Println("Input radiuses of two circles (r1; r2) ->")
	fmt.Scan(&r1,&r2)
	if r1+r2==math.Sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1))	{
		fmt.Print("One point crossing")
	}else{
		if r1+r2<math.Sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1))	{
			fmt.Print("no crossing points")
		}else{
			fmt.Print("Two points crossing")	
			
		}
	}
}
