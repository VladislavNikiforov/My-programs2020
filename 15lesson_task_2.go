/* 2. Заполнить квадратный массив следующим образом:
   a.      b.      c.      d.      e.      f.      g.      h.  
   12345   10000   10000   12345   01111   12345   12345   11111
   23451   01000   02000   21234   20111   23456   22345   12222
   34512   00100   00300   32123   22011   34567   33345   12333
   45123   00010   00040   43212   22201   45678   44445   12344
   51234   00001   00005   54321   22220   56789   55555   12345
*/
package main
import "fmt"
const n=5
type Array [n][n]int
func PrintArr(arr Array){
	for y:=0;y<n;y++{
		for x:=0;x<n;x++{
			fmt.Printf("%2d",arr[y][x])
		}
		fmt.Println()
	}
}
func main(){
	var arr Array
	var operation int
	fmt.Print("Choose operation (1-8)(подпункт задания) ->")
	fmt.Scan(&operation)
	switch operation{
		case 1: 
			row:=0
			for y:=0;y<n;y++{
				for x:=0;x<n-row;x++{
					arr[y][x]=x+row+1
				}
				numb:=1
				for x:=n-row;x<n;x++{
					arr[y][x]=numb
					numb++
				}
				row++
			}
			PrintArr(arr)
		case 2:
			for y:=0;y<n;y++{
				for x:=y;x<n;x++{
					if x==y{
						arr[y][x]=1
					}
				}
			}
			PrintArr(arr)
		case 3:
			for y:=0;y<n;y++{
				for x:=y;x<n;x++{
					if x==y{
						arr[y][x]=x+1
					}
				}
			}
			PrintArr(arr)
	    case 4:
			for y:=0;y<n;y++{
				k:=0
				for x:=0;x<=y;x++{
					arr[y][x]=y+1-k
					k++
				}
				l:=2
				for x:=k;x<n;x++{
					arr[y][x]=l
					l++
				}
			}
			PrintArr(arr)
		case 5:
			for y:=0;y<n;y++{
				for x:=0;x<y;x++{
					arr[y][x]=2
				}
				for x:=y+1;x<n;x++{
					arr[y][x]=1
				}	
			}
			PrintArr(arr)
		case 6:
			k:=0
			for y:=0;y<n;y++{
				k++
				for x:=0;x<n;x++{
					arr[y][x]=x+k
				}
			}
			PrintArr(arr)
		case 7:
			for y:=0;y<n;y++{
				for x:=0;x<y;x++{
					arr[y][x]=y+1
				}
				k:=0
				for x:=y;x<n;x++{
					arr[y][x]=y+1+k
					k++
				}
			}
			PrintArr(arr)
		case 8:
			for y:=0;y<n;y++{
				for x:=0;x<y;x++{
					arr[y][x]=x+1
				}
				for x:=y;x<n;x++{
					arr[y][x]=y+1
				}
			}
			PrintArr(arr)		
	}
}
