/*
 a. сколько в массиве положительных/отрицательных чисел
 b. сколько нулей

 * */
package main
import "fmt"
const MaxLength=10
type Array[MaxLength]int
func PosNeg(a Array){
	var pos,neg,null int
	for _,v:=range a{
		if v>0{
			pos++
		}else{
			if v<0 {
				neg++
			}else{
				null++
			}
		}
	}
	fmt.Printf("The amount of possitive numbers is %d\nTheamount of negative numbers is %d \nThe amount of nulls is %d ",pos,neg,null)
}
func main(){
	a:= Array{-1,-5,-3,-6,0,0,12,4,8,-22}
	fmt.Printf("We have Array{-1,-5,-3,-6,0,0,12,4,8,-22}\n")
	PosNeg(a)
}
