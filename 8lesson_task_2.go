//2. встречается ли в слайсе число Х
package main
import "fmt"
func Appearance(sl[]int,n int)bool{
	f:=false
	for i:=range sl{
		if sl[i]==n{
			f=true
		}
	}
	return f
}
func main() {
	var n int
	sl:=make([]int,5)
	
	fmt.Println("Input 5 slice elements")
	for i:=0;i<len(sl);i++{
		fmt.Scan(&sl[i])  
	}
	
	fmt.Println("Input number ->")
	fmt.Scan(&n)
	
	if Appearance(sl,n) {
		fmt.Printf("%d is in slice",n)
	}else{
		fmt.Printf("%d is not in slice",n)
	}
	
}
