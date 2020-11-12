package main
import "fmt"
func drawFigure(n int){
	for i:=1; i<n;i++{
		for j:=1; j<=n;j++{
			fmt.Print("x")
		}
	fmt.Println()
	}
}
func main(){
	drawFigure(4)
}
