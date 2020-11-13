//5. сколько в слайсе положительных, отрицательных чисел, нулей
package main
import "fmt"
func Sign(sl[]int)(int,int,int){
	pos:=0
	neg:=0
	null:=0
	for _,v:=range sl{
		if v>0{
			pos++
		}else{
			if v<0{
				neg++
			}else{
				null++
			}
		}
	}
	return pos,neg,null
}
func main (){
	sl:=make([]int,10)
	
	fmt.Print("Input 10 slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])  
	}
	pos,neg,null:=Sign(sl)
	fmt.Printf("\nSlice %d have:\n %d positive elements\n %d negative elements\n %d nulls",sl,pos,neg,null)
}
