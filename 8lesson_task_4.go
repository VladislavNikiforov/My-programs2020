//4. сколько раз в слайсе встречается максимум, на каком месте максимум встречается в первый раз
package main 
import "fmt"
func MaxPlace(sl[]int)(int,int){
	max:=sl[0]
	place:=0
	for i:=0;i<5 && i+1<5;i++{
		if max<sl[i+1]{
			max=sl[i+1]
			place=i+1
		}
	}
	return max,place
}
func main() {
	
	sl:=make([]int,5)
	
	fmt.Print("Input 5 slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])  
	}
	
	m,p:=MaxPlace(sl)
	fmt.Print("Maximum number; First place at which it appears: ",m,p)
}
