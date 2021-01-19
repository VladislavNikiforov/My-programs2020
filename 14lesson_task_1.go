/*1.    Бегать по экрану и закрашивать его:
    - туда-сюда под углом 45 градусов*/

package main
import ( 
	"fmt"
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
	fmt.Print(w,"  ",h)
	/*y:=0
	for y<h{
		for i:=0;i<w;i++{
			termbox.SetCell(i,y,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
		}
		y++
		for i:=w-1;i>=0;i--{
			termbox.SetCell(i,y,' ',termbox.ColorRed,termbox.ColorRed)
			termbox.Flush()
			time.Sleep(1 * time.Millisecond)
		}
		y++
	}
	*/
	time.Sleep(2*time.Second)
}
