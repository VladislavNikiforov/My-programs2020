/*
 b. на каком месте максимум/минимум встречается в первый/в последний раз
 a. сколько раз в массиве встречается максимум/минимум 
 */
package main
import "fmt"
const MaxLength=10
type Array[MaxLength]int
func CountMax(a Array) {
	for i:=range a{
		for j:=range a{
			if a[i]<a[j]{
				a[i],a[j]=a[j],a[i]
			}
		}
		
	}
	max:=a[MaxLength-1]
	maxC:=0
	min:=a[0]
	minC:=0
	for k:=MaxLength-1;k>=0;k--{
			if max==a[k]{
				maxC++
			}else{
				break
			}
	}
	
	for k:=0;k<MaxLength;k++{
		if min==a[k]{
			minC++
		}else{
			break
		}
	}
	maxFirst:=MaxLength-maxC
	minLast:=minC-1
	fmt.Printf("The programm outputs: sorted array %d\nMin numbers locates in  0-%d\nThe amount of Min numbers is %d\nMax numbers locates in %d - %d\nThe amount of Max numbers is %d",a,minLast,minC,maxFirst,MaxLength-1,maxC)
}
func main(){
	a:= Array{1,1,6,1,9,6,9,1,9,1}
	fmt.Println("We have Array{1,1,6,1,6,6,3,2,9,1}")
	CountMax(a)
}
