//12.Создать слайс с числами Фибоначчи, не превосходящими 1000000
package main

import "fmt"

func main(){
	
	sl:=make([]int,100)
	a,b,c:=1,0,0
	k:=0
	for i:=1; i+1<len(sl);i+=2 {
		c=b
		a+=c
		b+=a
		if a>1000000 || b>1000000{
			k=i
			break
		}
		sl[i]=a
		sl[i+1]=b
	}
	sl2:=sl[:k]
	fmt.Print(sl2)
}
