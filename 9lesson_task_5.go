// 5. Отсортировать слайс. 
// a). Выбором наименьшего.
// b). Пузырьком. 
// c). Шейкером. 
// Сделал через файл, чтобы потренироваться
package main

import "fmt"
import "os"
import "bufio"

func SelectionSort(sl[]int){
	var n = len(sl)
	for i:=0;i<n;i++{
		var k =i
		for j:=i;j<n;j++{
			if sl[j]<sl[k]{
				k=j
			}
		}
		sl[i],sl[k]=sl[k],sl[i]
	}
}

func Bubble(sl[]int){
	for i:=range sl{
		for j:=range sl{
			if sl[i]<sl[j]{
				sl[i],sl[j]=sl[j],sl[i]
			}
		}
	}
}

func Cocktail(sl[]int){
	n:=len(sl)-1
	for {
		f:=false
		for i:=0;i<n;i++{
			if sl[i]>sl[i+1]{
				sl[i+1],sl[i]=sl[i],sl[i+1]
				f=true
			}
		}
		if !f{
			return
		}
		f=false
		for i:=n-1;i>=0;i--{
			if sl[i]>sl[i+1]{
				sl[i+1],sl[i]=sl[i],sl[i+1]
				f=true
			}
		}
		if !f{
			return
		}
	}
}
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input95.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	defer fin.Close()
	in:=bufio.NewReader(fin)
	
	for {
		var c int
		_,err:=fmt.Fscan(in,&c)
		if err!=nil{
			break
		}
		sl=append(sl,c)
	}
	SelectionSort(sl) //наименьшее
	//Bubble(sl) //пузырь
	//Cocktail(sl) //шейкер
	fout,err:=os.Create("output95.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
		
	}
	defer fout.Close()
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	for i:=range sl{
		fmt.Fprintln(out, sl[i])
	}
}





