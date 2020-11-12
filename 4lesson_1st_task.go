package main 
import "fmt"
func main () {
	var a,b int = 1,1
	for a,b=1,1; a<=1000000 ||  b<=1000000; {
		fmt.Println(a)
		fmt.Println(b)
		S:=b
		a+=S
		b+=a
	}

}
