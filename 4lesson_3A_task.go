package main
import "fmt"
func main () {
	const 
	eps=1e-10
	var i int =2
	answer:=1.0
	x:=1.0/1.0
	for x>eps {
		answer+=x
		f:=1
		for j:=2; j<=i; j++ {
			f*=j
		}
		i++
		x=1/float64(f)
	}
fmt.Print(answer)		
		
}
