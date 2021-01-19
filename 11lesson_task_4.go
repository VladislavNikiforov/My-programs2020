//4. сколько в файле английских букв - и больших, и маленьких
package main
import "fmt"
import "os"
import "bufio"
func main(){
	f,err:=os.Open("input114.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	letters:=0
	
	for scanner.Scan(){
		str:=scanner.Text()
		sl:=[]rune(str)
		
		for _,v:=range sl{
			if v>='a' && v<='z' || v>='A' && v<='Z'{
				letters++
			}
		} 	
	}
	fmt.Printf("The amount of english letters in file is %d",letters)
}
