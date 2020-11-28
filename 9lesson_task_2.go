//2. Определить наибольший элемент файла,
// номер первого и последнего вхождения, 
//количество вхождений наибольшего числа в файл.

package main

import "fmt"

import "os"

import "bufio"

func Max(sl[]int)int{
	var max int
	max=sl[0]
	for i:=0;i+1<len(sl);i++{
		if max<sl[i]{
			max=sl[i]
		}
	}
	return max
}
func First(sl[]int,max int)int{
	var k int
	for i:=range sl{
		if sl[i]==max{
			k=i
			break
		}
	}
return k
}
func Last(sl[]int,max int)int{
	var k int
	for i:=range sl{
		if sl[i]==max{
			k=i
		}
	}
return k
}
func Count(sl[]int,max int)int{
	cnt:=0
	for i:=range sl{
		if sl[i]==max{
			cnt++
		}
	}
return cnt
}
func main(){
	var c int
	sl:=make([]int,0)
	fin,err:=os.Open("input92.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	in:=bufio.NewReader(fin)
	defer fin.Close()
	
	for {
		_,err:=fmt.Fscan(in,&c)
			if err!=nil{
				break
			}
			sl=append(sl,c)
	}
	fout,err:=os.Create("output92.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	defer fout.Close()
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	max:=Max(sl)
	fmt.Fprintf(out,"The maximum value in file is %d\nFirst time it appears at posotion %d\nLast - %d\nMaximum arrives %d times",Max(sl),First(sl,max),Last(sl,max),Count(sl,max))
}
