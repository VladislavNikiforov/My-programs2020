//1. сколько в файле символов
package main

import "fmt"
import "os"
import "bufio"

func main(){
	f,err:=os.Open("input111.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	scanner:=bufio.NewScanner(f)
	defer f.Close()
	var s int
	for scanner.Scan() {
		str:=scanner.Text()
		sl:=[]rune(str)
		s+=len(sl)
	}
	fmt.Print(s)
}
