/*
3. Лабиринт. Начальная позиция (начальное и конечное положение, стенки и свободные клетки)
 читаем из текстового файла. Цель, разумеется – выйти. Управление – стрелками. 
*/
package main
import (
	"fmt"
	"github.com/nsf/termbox-go"
	"bufio"
	"time"
	"os"
)


type Array [30][120]bool

func Wall(p,q int, arr Array)bool{ //если есть стена, выдает true
	if arr[q][p]{
		return arr[q][p]
	}else{ 
		return false
	}
}
func Win(m,n,y,x int){
	if n==y && m==x{
		s:="You are a winner!"
		slS:=[]rune(s)
		for i:=range slS{
			termbox.SetCell(x-5+i,y+1,slS[i],termbox.ColorWhite,termbox.ColorBlack)
			termbox.Flush()
		}
		time.Sleep(2*time.Second)
		termbox.Close()
		os.Exit(1)
	}
}
func main(){
	var arr Array
	err:=termbox.Init()
	if err!=nil{
		return
	}
	defer termbox.Close()
	x,y:=0,0
	w,h:=termbox.Size()
	fin,err:=os.Open("input153_2.txt")
	if err!=nil{
		fmt.Print("file read error")
		
		return
	}
	defer fin.Close()
	in:=bufio.NewScanner(fin)
	//скан документа и рисовалка лабиринта
	var n,m int //координаты финиша
	for in.Scan(){
		str:=in.Text()
		sl:=[]rune(str)
		for i,v:=range sl{
			switch v{
			case '*':
				termbox.SetCell(x+i,y,' ',termbox.ColorYellow,termbox.ColorYellow)
				termbox.Flush()
			case '?':
				n,m=y,i
				termbox.SetCell(x+i,y,' ',termbox.ColorGreen,termbox.ColorGreen)
				termbox.Flush()
			default:
				termbox.SetCell(x+i,y,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
			}
		}
	//заполнение массива положением стен
		for j,v:=range sl{
			if v=='*'{
				arr[y][j]=true
			}else{
				arr[y][j]=false
			}
		}
		y++
	}
	//появление игрока
	x,y=0,0
	termbox.SetCell(x,y,'^',termbox.ColorGreen,termbox.ColorBlack)
	termbox.Flush()
	//управление
	for {
		ev:=termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyArrowDown:
			p,q:=x,y+1
			if y+1<h && !Wall(p,q,arr){
				termbox.SetCell(x,y,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
				termbox.SetCell(p,q,'v',termbox.ColorGreen,termbox.ColorBlack)
				termbox.Flush()
				y++
				Win(m,n,y,x)
			}
		case termbox.KeyArrowUp:
			p,q:=x,y-1
			if y>0 && !Wall(p,q,arr){
				termbox.SetCell(x,y,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
				termbox.SetCell(p,q,'^',termbox.ColorGreen,termbox.ColorBlack)
				termbox.Flush()
				y--
				Win(m,n,y,x)
			}
			
		case termbox.KeyArrowRight:
			p,q:=x+1,y
			if x<w-1 && !Wall(p,q,arr){//булеан функция на проверку 1 или нуля из массива 
				termbox.SetCell(x,y,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
				termbox.SetCell(p,q,'>',termbox.ColorGreen,termbox.ColorBlack)
				termbox.Flush()
				x++
				Win(m,n,y,x)
			}
		case termbox.KeyArrowLeft:
			p,q:=x-1,y
			if x>0 && !Wall(p,q,arr){
				termbox.SetCell(x,y,' ',termbox.ColorBlack,termbox.ColorBlack)
				termbox.Flush()
				termbox.SetCell(p,q,'<',termbox.ColorGreen,termbox.ColorBlack)
				termbox.Flush()
				x--
				Win(m,n,y,x)
			}
		}
	}
	
	
}
