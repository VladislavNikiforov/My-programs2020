package main
import "fmt"
func main () {
	var a,b int
	fmt.Scan(&a,&b)
	if a<b {
		for a<=b {
			fmt.Println(a)
			a++
		}
	}else{
		if a>b {
			for a>=b {
				fmt.Println(a)
				a--
			}
		}else{
			fmt.Print(a)
		}
	}
}
