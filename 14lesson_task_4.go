//4. Гистограмма – обычная: вертикальные столбцы, идущие снизу вверх. 
//Может быть на столбцах (столбцы шириной в несколько позиций) писать сотвествующие числа. Данные читаем из файла. 
package main
import (
	"fmt"
	"os"
	"bufio"
	"time"
	"github.com/nsf/termbox-go"
)
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input144.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	in:=bufio.NewReader(fin)
	defer fin.Close()
	
	err2:=termbox.Init()
	if err2!=nil{
		fmt.Print("Can't open termbox")
		return
	}
	defer termbox.Close()
	w,h:=termbox.Size()
	
	for {
		var c int
		_,err:=fmt.Fscan(in,&c)
		if err!=nil{
			break
		}
		if c>=0 && c<=h && len(sl)<w{
			sl=append(sl,c)
		}
	}
	
	
	max:=sl[0]
	for i:=range sl{
		if max<sl[i]{
			max=sl[i]
		}
	}
	maxInd:=max
	for i:=0; i<max; i++{
		for j:=range sl{
			if sl[j]>maxInd{
				termbox.SetCell(j,i,' ',termbox.ColorBlack,termbox.ColorRed)
				termbox.Flush()
					
			}else{
				termbox.SetCell(j,i,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
			
			}
		}
		maxInd--
	}
	time.Sleep(5*time.Second)
}






