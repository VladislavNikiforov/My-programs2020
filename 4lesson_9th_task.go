/*package main 
import "fmt"
func main () {
//9. Имеется последовательность: 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, ... Какое число находится на N-ом месте.
	var x int
	var i int =1
	fmt.Println("Input serial number, programm calculates the value")
	fmt.Scan(&x)
		for x>0 {
			x=x-i
			i++
			if x<=0 {
				i--
			}
		}
	fmt.Printf("The numbers value is %d",i)
		
}
*/
//ivan version
package main
import "fmt"
import "math"
func main () {
    var a, b float64
    fmt.Println("Write number")
    fmt.Scanln(&a)

    if a>0 {
         b = math.Sqrt(2*a) + 1/2
     }
    b=math.Round(b)
    fmt.Println(b)
}
