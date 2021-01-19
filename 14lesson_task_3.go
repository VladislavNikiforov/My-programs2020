//3. Ввод пароля – жмём буквы, а на экране появляются звёздочки
package main
import "github.com/nsf/termbox-go"
import "time"
func main(){
	var pass string
	err:=termbox.Init()
	if err!=nil{
		return
	}
	defer termbox.Close()
	x:=0
	for {
		ev:=termbox.PollEvent()
		if ev.Key==termbox.KeyCtrlC{
			break
		}
		if ev.Key==termbox.KeyEnter{
			sl:=[]rune(pass)
			for i:=range sl{
				termbox.SetCell(i,1,sl[i],termbox.ColorRed,termbox.ColorYellow)
				termbox.Flush()
				x++
			}
			time.Sleep(3*time.Second)
		}else{
			if ev.Type==termbox.EventKey{
				termbox.SetCell(x,0,'*',termbox.ColorRed,termbox.ColorYellow)
				termbox.Flush()
				x++
				pass+=string(ev.Ch)
			}
		
		}
	}
}
