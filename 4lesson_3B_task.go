package main
import "fmt"
func main () {
	const 
	eps=1e-3
	var f int =1
	answer:=1.0
	x:=1.0/1.0
	for x>eps {
		f*=3
		x=1/float64(f)
		answer+=x
	}
fmt.Print(answer)		
		
}
