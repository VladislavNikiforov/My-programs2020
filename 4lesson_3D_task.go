package main
import "fmt"
func main () {
	const eps=1e-8

	var i int =2
	answer:=0.0
	x:=1.0
	p:=1
	for x>eps {
		f:=p
		p=i
		f=f*i*(i+1)
		
		x=1/float64(f)
		fmt.Println(x)
		answer+=x
		i++
	}
		fmt.Print(answer)
}
