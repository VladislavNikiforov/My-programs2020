//6. найти место последнего ненулевого элемента слайса
package main

import "fmt"

func LastNot0Place(sl[]int)(int){
	place0:=0
	for i,v:=range sl{
		if sl[0]==0{
			place0=-1
			break
		}
		if v==0{
			place0=i
		}
		if i==len(sl)-1 && place0==0{
			place0=0
		}
	}
	return place0
}

func main(){
	sl:=make([]int,10)
	
	fmt.Println("Input 10 slice elements")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	
	if LastNot0Place(sl)>0{
		fmt.Printf("The last not-null element is on place %d",LastNot0Place(sl)-1)
	}
	if LastNot0Place(sl)==0{
		 fmt.Print("The slice doesn't have nulls")
	 }
	if LastNot0Place(sl)==-1{
		fmt.Print("First slice element is 0")
	}
}
