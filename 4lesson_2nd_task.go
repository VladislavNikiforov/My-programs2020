package main
import "fmt"
func main () {
	var a int
	fmt.Scan(&a)
	for i:=2; i!=a; i++ {
		if a%i==0 {
			fmt.Print("common")
			break
		}else{ 
			if i==a-1 {
				fmt.Print("Prime")
			}
		}
	}
}
