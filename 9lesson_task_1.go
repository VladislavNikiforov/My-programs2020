//1. Определить номер первого и последнего вхождения,
//количество вхождений данного числа в файл.

package main

import "fmt"

import "os"

import "bufio"

func First(sl[]int,n int)int{
	var k int
	for i:=range sl{
		if sl[i]==n{
			k=i
			break
		}
	}
	return k
}

func Last(sl[]int,n int)int{
	var k int
	for i:=range sl{
		if sl[i]==n{
			k=i
		}
	}
	return k
}
func Count(sl[]int,n int)int{
	cnt:=0
	for i:=range sl{
		if sl[i]==n{
			cnt++
		}
	}
return cnt
}

func main(){
	var n,c int
	sl:=make([]int,0)
	fin,err:=os.Open("input91.txt")
	if err!=nil{
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
	
	fmt.Print("Input value -> ")
	fmt.Scan(&n)
	
	fout,err:=os.Create("output91.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	defer fout.Close()
	
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	
	fmt.Fprintf(out,"First time value arrives in file at %d place\nLast time - %d\nIt arrives %d times",First(sl,n),Last(sl,n),Count(sl,n))
	
}
