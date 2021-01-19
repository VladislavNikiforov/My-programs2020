//2. сколько в файле слов
package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main(){
	f,err:=os.Open("input112.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	scanner:=bufio.NewScanner(f)
	defer f.Close()
	
	words:=0
	
	for scanner.Scan(){
		str:=scanner.Text()
	
		Spaces:=strings.Split(str," ")
		for i:=0;i<len(Spaces);{
			if Spaces[i]==""{
				Spaces=append(Spaces[:i],Spaces[i+1:]...)
			}else{
				i++
			}
		}
		words+=len(Spaces)
	}
	fmt.Print(words)
}
