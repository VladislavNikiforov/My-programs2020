/*5. Точный ли квадрат. function PerfectSquare(n uint) (ok uint, root2 int) Если число n является точным квадратом, то возвращает ok=1,
 root2 - квадратный корень из n; если не является точным квадратом, то ok = 0, а root2 не определяем*/
package main
import "fmt"
import "math"
func PerfectSquare(n uint) (ok uint, root2 uint){
	var x float64
	var i uint=1
	for ; i*i<n ;i++ {
	}
		if i*i==n {
			x=math.Sqrt(float64(n))
			return 1,uint(x)
		}else{
			 return 0,0
		 }
}
func main () {
	var a uint
	fmt.Scanln(&a)
	fmt.Print(PerfectSquare(a))
}
