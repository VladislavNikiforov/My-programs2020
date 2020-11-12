package main
import "fmt"
func drawFig(n int){
	for i:=1; i<=n;i++{
		fmt.Print("x")
	}
	fmt.Println()
	for j:=1; j<=2;j++{
		fmt.Print("x")
		for l:=1;l<=4;l++{
			fmt.Print("+")
		}
		fmt.Print("x")
		fmt.Println()
	}
	for i:=1; i<=n;i++{
		fmt.Print("x")
	}
		
}
func main (){
	drawFig(6)
}
