/*1.    Бегать по экрану и закрашивать его:
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
	y:=0
	x:=0
	termbox.SetCell(x,y,' ',termbox.ColorRed,termbox.ColorRed)
	termbox.Flush()
	for { //программа не заканчивается, так как нет условия для этого. при условии x<w && y<h - не работает
		
		x++
		termbox.SetCell(x,y,' ',termbox.ColorRed,termbox.ColorRed)
		termbox.Flush()
		time.Sleep(1 * time.Millisecond)
		for ; x>=0 && x<w && y<h; x--{
			
			termbox.SetCell(x,y,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
			y++
		}
		y++
		termbox.SetCell(x,y,' ',termbox.ColorRed,termbox.ColorRed)
		termbox.Flush()
		time.Sleep(1 * time.Millisecond)
		
		for ; y>=0 && y<h && x<w; y--{
			
			termbox.SetCell(x,y,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
			x++
		}
		
	}
	
	//time.Sleep(3*time.Second)
}
