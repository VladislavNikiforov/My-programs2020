//8. Делится ли одно число на другое.
package main
import "fmt"
func div(a,b int)int{ 
	var x int
	if a==0||b==0 {
		return 0
	}
	if a%b==0|| b%a==0 {
		x=1
	}
	if a%b!=0 && b%a!=0 {
		x=0
	}
	return x
}
func main() {
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Print(div(a,b))
}
