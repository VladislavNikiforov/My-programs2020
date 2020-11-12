package main
import (
"fmt"
"math"
)
func main() {
var x1,x2,y1,y2,ax,ay,a float64
fmt.Print("Input first point coordinates -> ")
fmt.Scanln(&x1,&y1)
fmt.Print("Input second point coordinates -> ")
fmt.Scanln(&x2,&y2)
ax=x2-x1
ay=y2-y1
a=math.Sqrt(math.Pow(ax,2)+math.Pow(ay,2))
fmt.Print("The lenght of a vetor is ",a)
}
