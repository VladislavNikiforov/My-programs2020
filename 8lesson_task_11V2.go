//Есть два возрастающих слайса, склеить их в один, чтобы он тоже возрастал.

package main

import "fmt" 

func main(){
	sl:=make([]int,5)
	sl2:=make([]int,5)
	sl3:=make([]int,0)
	var k,l int
	
	fmt.Println("Input 5 growing slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	fmt.Println("Input another 5 growing slice elements ->")
	for i:=range sl2{
		fmt.Scan(&sl2[i])
	}
	
	for j:=0;j<=10 && k!=5 && l!=5; j++{
		if sl[k]<sl2[l]{
			sl3=append(sl3,sl[k])
			k++
		}
		if sl[k]>sl2[l]{
			sl3=append(sl3,sl2[l])
			l++
		}
		if sl[k]==sl2[l]{
			sl3=append(sl3,sl[k])
			sl3=append(sl3,sl2[l])
		}
	}
	fmt.Print(sl3)
}
