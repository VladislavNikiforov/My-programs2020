 /*1. Заполнить прямоугольный массив следующим образом:
 a.   01010101    b. 01010101  c. 11111111  d. 10101010    
      10101010       11111111     10000001     02020202     
      01010101       01010101     10000001     30303030    
      10101010       11111111     11111111     04040404
*/
package main
import (
	"fmt"
)
const (
	h=4
	w=8
)
type Array [4][8]int

func PrintArr(arr Array){
	for y:=0;y<h;y++{
			for x:=0;x<w;x++{
				fmt.Printf("%2d",arr[y][x])
			}
			fmt.Println()
		}
}
func main(){
	var arr[h][w]int
	var num int
	fmt.Println("Choose operation 1-4")
	fmt.Scan(&num)
	
	switch {
	case num==1:
		for y:=0;y<h;y++{
			for x:=0;x<w;x++{
				arr[y][x]=(x+y)%2
			}
		}
		PrintArr(arr)
	
	case num==2:
		y:=0
		for {
			for x:=0;x<w;x++{
				if x%2!=0{
					arr[y][x]=1
				}
			}
			y++
			if y==h{
				break
			}
			for x:=0;x<w;x++{
				arr[y][x]=1
			}
			y++
			if y==h{
				break
			}
		}
		PrintArr(arr)
	
	case num==3:
		y:=0
		for y<h{
			switch {
			case y==0 || y==h-1:
				for x:=0;x<w;x++{
					arr[y][x]=1
				}
				y++
			default:
				for x:=0;x<w;x++{
					if x==0 || x==w-1 {
						arr[y][x]=1
					}else{
						arr[y][x]=0
					}
				}
			y++
			}
		}
		PrintArr(arr)
			
	case num==4:
		num:=1
		for y:=0;y<h;y++{
			for x:=0;x<w;x++{
				switch {
				case (y+x)%2==0: 
					arr[y][x]=num
				default: 
					arr[y][x]=0
				}
				
			}
			num++
		}
		PrintArr(arr)
	}
	
}






