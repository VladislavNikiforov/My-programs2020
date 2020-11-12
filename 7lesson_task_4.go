/*
4. найти место последнего ненулевого элемента массива (len, если на месте (len-1) стоит не 0)
*/
package main
import "fmt"
const MaxLength=10
type Array[MaxLength]int
func len0(a Array)int{
	len0:=0
	for i,v:=range a {
		if v==0{
			len0=i
			break
		}
	}
	return len0
}
func main(){ 
	a:=Array{1,2,5,6,4,0,7,8,9,11}
	fmt.Printf("We have Array %d\nThe last not null number index is %d",a,len0(a))
}
