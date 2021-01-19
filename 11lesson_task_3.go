//3. сколько в файле цифр
package main

import "fmt"
import "os"
import "bufio"

func main(){
	f,err:=os.Open("input113.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	numbers:=0
	for scanner.Scan(){
		str:=scanner.Text()
		sl:=[]rune(str)
		for _,v:=range sl{
			if v>='0' && v<='9'{
				numbers++
			}
		}
	}
	fmt.Printf("The amount of numbers in file is %d",numbers)
}


		
