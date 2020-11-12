/*
  a. встречается ли в массиве число Х (это bool)
  b. сколько раз число Х входит в массив (это int)
*/
package main
import "fmt"
const MaxLength=10
type Array [MaxLength] int
func Check(a Array,x int)bool {
	f:=false
	for _,v:=range a{
		if v==x {
			f=true
		}
	}
	return f
}
func Count(a Array,x int)int{
	k:=0
	for _,v:=range a{
		if v==x{
			k++
		}
	}
	return k
}
func main(){
	a:= Array{1,6,12,6,11,11,8,7,6,4}
	var x,o int
	fmt.Print("Input the value ->")
	fmt.Scanln(&x)
	fmt.Print("Choose the operation: 1 or 2 ->")
	fmt.Scanln(&o)
	if o==1{
		fmt.Print(Check(a,x))
	}else{ 
		if o==2{
			fmt.Print(Count(a,x), " times")
		}else{
			fmt.Print("Error")
		}
	}
}
