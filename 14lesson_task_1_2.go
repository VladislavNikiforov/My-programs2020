/*1.    Бегать по экрану и закрашивать его:
    - вправо-влево
    - вниз-вверх
    - туда-сюда под углом 45 градусов*/

package main
import ( 
	"time"
	"github.com/nsf/termbox-go"
)
func main(){
	err:=termbox.Init()
	if err!=nil{
		return
	}
	defer termbox.Close()
	w,h:=termbox.Size()
	x:=0
	for x<w{
		for i:=0;i<h;i++{
			termbox.SetCell(x,i,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
		}
		x++
		for i:=w-1;i>=0;i--{
			termbox.SetCell(x,i,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
		}
		x++
	}
	
	time.Sleep(2*time.Second)
}
