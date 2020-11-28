//5. возвращает место K-го вхождения данного символа в строку
package main
import "fmt"
func Positions(s string, x string, n int)int{
	sl:=[]rune(s)
	slX:=[]rune(x)
	var place,cnt int
	for i,v:=range sl{
		if v==slX[0]{
			cnt++
			if cnt==n{
				place=i
				break
			}
		} 
	}
	if cnt<n{
		fmt.Println("Incorrect number or letter")
		place=-1
	}
	return place
}
func main(){
	s:="repeating letters"
	var x string
	var n int
	fmt.Printf("We have word {%s},\nInput a letter and a time(раз) it appears\nProgram outputs that letter position ->",s)
	fmt.Scan(&x,&n)
	fmt.Printf("Position is %d",Positions(s,x,n))	
}
