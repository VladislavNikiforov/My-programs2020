package main
import "fmt"
func main () {
	const eps=1e-3
	var i int =2
	answer:=0.0
	x:=1.0
	for x>eps {
		var p int=1
		f:=p
		f*=i
		p=i
		x=1/float64(f)
		fmt.Println(x)
		answer+=x
		i++
	}
		fmt.Print(answer)
}
